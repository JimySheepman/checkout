package main

import (
	"checkout-case/config"
	"checkout-case/file"
	"checkout-case/handler"
	"checkout-case/pkg/logger"
	"checkout-case/repository"
	"checkout-case/rest"
	"checkout-case/services"
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ServerTypeREST = "rest"
	ServerTypeFILE = "file"
)

type cartRepository interface {
	Create() error
}

type service struct {
	restServer *rest.RestServer
	fileServer *file.FileServer
}

func (s *service) InitCart(cartRepo cartRepository) error {
	if err := cartRepo.Create(); err != nil {
		return fmt.Errorf("initil cart error: %w", err)
	}

	return nil
}

func (s *service) Start() {
	l := logger.GetLogger().Sugar()

	var (
		err         error
		errRestChan = make(chan error, 10)
		errFileChan = make(chan error, 10)
	)

	cartRepo := repository.NewCartRepository()

	if err := s.InitCart(cartRepo); err != nil {
		l.Error(err)
		return
	}
	l.Info("init cart successfully created")

	promotionService := services.NewPromotionService()
	cartService := services.NewCartService(cartRepo, promotionService)

	restHandler := handler.NewRestHandler(cartService)
	fileHandler := handler.NewFileHandler(cartService)

	s.restServer = rest.NewRestServer(restHandler)
	s.fileServer = file.NewFileServer(fileHandler)

	switch config.Config.Server.ServerType {
	case ServerTypeREST:
		err = s.restServer.Start(errRestChan)
		if err != nil {
			l.Errorf("rest server init error: %v", err)
			return
		}
		l.Info("rest server successfully started")
	case ServerTypeFILE:
		err = s.fileServer.Start(errFileChan)
		if err != nil {
			l.Errorf("file server init error: %v", err)
			return
		}
		l.Info("file server successfully started")
	default:
		l.Errorf("unknown server type failure")
		return
	}

	s.waitSignal(errFileChan, errRestChan)
}

func (s *service) waitSignal(errFileChan chan error, errRestChan chan error) {
	l := logger.GetLogger().Sugar()

	signalChan := make(chan os.Signal, 1)
	// syscall package and SIGTERM may be removed for Windows
	// os.Interrupt (SIGINT) and os.Kill (SIGKILL) are only signal values guaranteed to be present on all systems
	// But os.Kill cannot be trapped by a program
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case err := <-errFileChan:
			if err != nil {
				l.Errorf("File Server error: %+v", err)
			}
			l.Info("Shutdown completed...")

			os.Exit(0)
		case err := <-errRestChan:
			if err != nil {
				l.Errorf("Rest Server error: %+v", err)
			}
			l.Info("Shutdown completed...")

			os.Exit(0)
		case <-signalChan:
			l.Info("Signal received, shutting down...")
			s.Stop()
		}
	}
}

func (s *service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.restServer.GracefulShutdown(ctx)
	logger.GetLogger().Sugar().Info("rest server graceful shutdown is done")
}

func InitLogger(c config.LoggerConfig) *zap.Logger {
	var loggerOpts []logger.LoggerOption

	loggerOpts = append(loggerOpts, logger.WithIO(os.Stdout, c.LogLevel, c.EnvironmentType, c.LogEncoding))

	if c.GraylogAddr != "" {
		loggerOpts = append(loggerOpts, logger.WithGraylogViaUDP(c.LogLevel, c.GraylogAddr))
	}

	logger.InitLogger(c.AppName, loggerOpts...)

	return logger.GetLogger()
}
