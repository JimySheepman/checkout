package file

import (
	"bufio"
	"checkout-case/internal/core/models"
	"checkout-case/internal/core/port"
	"checkout-case/pkg/config"
	"checkout-case/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
)

const (
	unixEOF                 = "\n"
	addItemCommand          = "addItem"
	addVasItemToItemCommand = "addVasItemToItem"
	removeItemCommand       = "removeItem"
	resetCartCommand        = "resetCart"
	displayCartCommand      = "displayCart"

	loggerCommandIdKey = "commandId"
)

type CommandHandler func(ctx context.Context, input string) (string, error)

type FileServer struct {
	fileHandlerClient port.FileHandlerClient
}

func NewFileServer(fileHandlerClient port.FileHandlerClient) *FileServer {
	return &FileServer{
		fileHandlerClient: fileHandlerClient,
	}
}

func (s *FileServer) Start(errChan chan error) error {
	_, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		errChan <- s.execute()
	}()

	return nil
}

func (s *FileServer) execute() error {
	l := logger.GetLogger()
	l.Info("started file server execute")

	lines, err := s.read()
	if err != nil {
		l.Sugar().Errorf("reading all commands error: %v", err)
		return err
	}
	l.Info("file server read all commands")

	for i, line := range lines {
		l.With(
			zap.Int64(loggerCommandIdKey, int64(i+1)),
		)
		ctx := logger.WithCtx(context.Background(), l)

		l.Info("started command unmarshal")
		cmd, err := s.commandUnmarshal(ctx, line)
		if err != nil {
			if err := s.writeFailureResponse(ctx, err.Error()); err != nil {
				l.Error("writing an errored command response to a file failed.")
			}
			continue
		}
		l.Info("done command unmarshal")

		l.Info("started find command handler")
		commandHandler, err := s.findCommandHandler(cmd)
		if err != nil {
			if err := s.writeFailureResponse(ctx, err.Error()); err != nil {
				l.Error("writing a non-handler command response to a file failed.")
			}
			continue
		}
		l.Info("done find command handler")

		l.Info("started command handler")
		resp, err := commandHandler(context.TODO(), line)
		if err != nil {
			if err := s.writeFailureResponse(ctx, err.Error()); err != nil {
				l.Error("writing an errored command handler response to a file failed.")
			}
			continue
		}
		l.Info("done command handler")

		l.Info("started write succeed response")
		if err := s.writeSucceedResponse(ctx, resp); err != nil {
			l.Error("writing an succeed command handler response to a file failed.")
			continue
		}
		l.Info("done write succeed response")
	}

	return nil
}

func (s *FileServer) read() ([]string, error) {
	l := logger.GetLogger().Sugar()
	l.Info("started file read")

	f, err := os.Open(config.Cfg.Server.FileServer.InputPath)
	if err != nil {
		l.Errorf("file open error: %v", err)
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cmd []string
	for scanner.Scan() {
		cmd = append(cmd, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		l.Errorf("file scan error: %v", err)
		return nil, err
	}
	l.Info("done file read")

	return cmd, nil
}

func (s *FileServer) writeSucceedResponse(ctx context.Context, input string) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("started write succeed response")

	f, err := os.OpenFile(config.Cfg.Server.FileServer.OutputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.Errorw(
			fmt.Sprintf("file open error: %v", err),
			"outputPath", config.Cfg.Server.FileServer.OutputPath,
		)
		return err
	}
	defer f.Close()

	out := []byte(unixEOF + populateSucceedGenericCommandResponseToString(input))
	if _, err := f.Write(out); err != nil {
		l.Errorf("write file error: %v", err)
		return err
	}
	l.Info("done write succeed response")

	return nil
}

func populateSucceedGenericCommandResponseToString(input string) string {
	resp := &models.GenericCommandResponse{
		Result:  true,
		Message: input,
	}

	bResp, _ := json.Marshal(resp)

	return string(bResp)
}

func (s *FileServer) writeFailureResponse(ctx context.Context, input string) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("started write failure response")

	f, err := os.OpenFile(config.Cfg.Server.FileServer.OutputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer f.Close()

	out := []byte(unixEOF + populateFailureGenericCommandResponseToString(input))
	if _, err = f.Write(out); err != nil {
		return fmt.Errorf("%w", err)
	}
	l.Info("done write failure response")

	return nil
}

func populateFailureGenericCommandResponseToString(input string) string {
	resp := &models.GenericCommandResponse{
		Result:  false,
		Message: input,
	}

	bResp, _ := json.Marshal(resp)

	return string(bResp)
}

func (s *FileServer) commandUnmarshal(ctx context.Context, text string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	cmd := &models.CommandRequest{}
	if err := json.Unmarshal([]byte(text), cmd); err != nil {
		l.Error("Unmarshal error")
		return "", err
	}

	if cmd.Command == "" {
		l.Error("command empty")
		return "", fmt.Errorf("command empty")
	}

	return cmd.Command, nil
}

func (s *FileServer) findCommandHandler(cmd string) (CommandHandler, error) {
	switch cmd {
	case addItemCommand:
		return s.fileHandlerClient.AddItemHandler, nil
	case addVasItemToItemCommand:
		return s.fileHandlerClient.AddVasItemToItemHandler, nil
	case removeItemCommand:
		return s.fileHandlerClient.RemoveItemHandler, nil
	case resetCartCommand:
		return s.fileHandlerClient.ResetCartHandler, nil
	case displayCartCommand:
		return s.fileHandlerClient.DisplayCartHandler, nil
	default:
		return nil, fmt.Errorf("command is not found")
	}
}
