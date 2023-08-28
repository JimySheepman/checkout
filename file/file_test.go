//go:generate mockgen -destination=./mocks/file_mock.go -source=file.go checkout-case/file

package file

import (
	"checkout-case/config"
	mock_file "checkout-case/file/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var ErrTest = errors.New("test error")

type fileServerMocks struct {
	mockFileHandlerClient *mock_file.MockfileHandlerClient
}

func setupFileServerTest(t *testing.T) (context.Context, *fileServerMocks, *FileServer) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &fileServerMocks{
		mockFileHandlerClient: mock_file.NewMockfileHandlerClient(ctrl),
	}

	srv := NewFileServer(mocks.mockFileHandlerClient)

	return ctx, mocks, srv
}

func TestFileServer_execute(t *testing.T) {
	_, mocks, srv := setupFileServerTest(t)

	cfg, err := config.Load()
	require.Nil(t, err)

	tests := []struct {
		name         string
		isError      bool
		expectations func()
	}{
		{
			name:    "s.read failure",
			isError: true,
			expectations: func() {
				cfg.Server.FileServer.InputPath = ""
			},
		},
		{
			name:    "writeFailureResponse failure",
			isError: false,
			expectations: func() {
				cfg.Server.FileServer.InputPath = "../input-test.txt"
				cfg.Server.FileServer.OutputPath = ""

				mocks.mockFileHandlerClient.EXPECT().AddItemHandler(gomock.Any(), gomock.Any()).Return("", ErrTest)
			},
		},
		{
			name:    "writeSucceedResponse failure",
			isError: false,
			expectations: func() {
				cfg.Server.FileServer.InputPath = "../input-test.txt"
				cfg.Server.FileServer.OutputPath = ""

				mocks.mockFileHandlerClient.EXPECT().AddItemHandler(gomock.Any(), gomock.Any()).Return("", nil)
			},
		},
		{
			name:    "execute succeed",
			isError: false,
			expectations: func() {
				cfg.Server.FileServer.InputPath = "../input-test.txt"
				cfg.Server.FileServer.OutputPath = "../output-test.txt"

				mocks.mockFileHandlerClient.EXPECT().AddItemHandler(gomock.Any(), gomock.Any()).Return("", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.execute()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestFileServer_read(t *testing.T) {
	_, _, srv := setupFileServerTest(t)

	cfg, err := config.Load()
	require.Nil(t, err)

	tests := []struct {
		name         string
		expected     []string
		isError      bool
		expectations func()
	}{
		{
			name:     "OpenFile failure",
			expected: nil,
			isError:  true,
			expectations: func() {
				cfg.Server.FileServer.InputPath = ""
			},
		},
		{
			name:     "read succeed",
			expected: []string{`{"command":"addItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`, `{"commant":"addItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`},
			isError:  false,
			expectations: func() {
				cfg.Server.FileServer.InputPath = "../input-test.txt"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.read()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileServer_writeSucceedResponse(t *testing.T) {
	ctx, _, srv := setupFileServerTest(t)

	cfg, err := config.Load()
	require.Nil(t, err)

	tests := []struct {
		name         string
		input        string
		isError      bool
		expectations func()
	}{
		{
			name:    "OpenFile failure",
			input:   "",
			isError: true,
			expectations: func() {
				cfg.Server.FileServer.OutputPath = ""
			},
		},
		{
			name:    "writeSucceedResponse succeed",
			input:   "test-input",
			isError: false,
			expectations: func() {
				cfg.Server.FileServer.OutputPath = "../output-test.txt"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.writeSucceedResponse(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestFileServer_writeFailureResponse(t *testing.T) {
	ctx, _, srv := setupFileServerTest(t)

	cfg, err := config.Load()
	require.Nil(t, err)

	tests := []struct {
		name         string
		input        string
		isError      bool
		expectations func()
	}{
		{
			name:    "OpenFile failure",
			input:   "",
			isError: true,
			expectations: func() {
				cfg.Server.FileServer.OutputPath = ""
			},
		},
		{
			name:    "writeFailureResponse succeed",
			input:   "test-input",
			isError: false,
			expectations: func() {
				cfg.Server.FileServer.OutputPath = "../output-test.txt"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.writeFailureResponse(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestFileServer_commandUnmarshal(t *testing.T) {
	ctx, _, srv := setupFileServerTest(t)

	tests := []struct {
		name     string
		text     string
		expected string
		isError  bool
	}{
		{
			name:     "json.Unmarshal failure",
			text:     "test-value",
			expected: "",
			isError:  true,
		},
		{
			name:     "command is empty failure",
			text:     `{"command":"","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: "",
			isError:  true,
		},
		{
			name:     "commandUnmarshal succeed",
			text:     `{"command":"addItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: addItemCommand,
			isError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := srv.commandUnmarshal(ctx, tt.text)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileServer_findCommandHandler(t *testing.T) {
	_, _, srv := setupFileServerTest(t)

	tests := []struct {
		name    string
		cmd     string
		isError bool
	}{
		{
			name:    "add item command succeed",
			cmd:     addItemCommand,
			isError: false,
		},
		{
			name:    "add vasItem to tem command succeed",
			cmd:     addVasItemToItemCommand,
			isError: false,
		},
		{
			name:    "remove item command succeed",
			cmd:     removeItemCommand,
			isError: false,
		},
		{
			name:    "reset cart command succeed",
			cmd:     resetCartCommand,
			isError: false,
		},
		{
			name:    "display cart command succeed",
			cmd:     displayCartCommand,
			isError: false,
		},
		{
			name:    "unknown command failure",
			cmd:     "test",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := srv.findCommandHandler(tt.cmd)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
