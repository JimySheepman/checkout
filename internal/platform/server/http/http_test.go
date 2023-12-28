//go:generate mockgen -destination=../../../../mocks/platform/http_mock.go -source=http.go

package http

import (
	mock_http "checkout-case/mocks/platform"
	"checkout-case/pkg/config"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type restServerMocks struct {
	mockRestHandlerClient *mock_http.MockrestHandlerClient
}

func setupRestServerTest(t *testing.T) (context.Context, *restServerMocks, *RestServer) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &restServerMocks{
		mockRestHandlerClient: mock_http.NewMockrestHandlerClient(ctrl),
	}

	srv := NewRestServer(mocks.mockRestHandlerClient)

	return ctx, mocks, srv
}

func TestRestServer_Start(t *testing.T) {
	ctx, mocks, srv := setupRestServerTest(t)

	err := config.LoadConfig()
	require.Nil(t, err)

	config.Cfg.Server.RestServer.PprofEnable = 1
	config.Cfg.Server.RestServer.Addr = ":9090"

	tests := []struct {
		name         string
		expectations func()
	}{
		{
			name: "healthcheck succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().HealthCheckHandler(gomock.Any()).Return(nil)
			},
		},
		{
			name: "AddItemHandler succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().AddItemHandler(gomock.Any()).Return(nil)
			},
		},
		{
			name: "AddVasItemToItemHandler succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().AddVasItemToItemHandler(gomock.Any()).Return(nil)
			},
		},
		{
			name: "RemoveItemHandler succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().RemoveItemHandler(gomock.Any()).Return(nil)
			},
		},
		{
			name: "ResetCartHandler succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().ResetCartHandler(gomock.Any()).Return(nil)
			},
		},
		{
			name: "DisplayCartHandler succeed",
			expectations: func() {
				mocks.mockRestHandlerClient.EXPECT().DisplayCartHandler(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errC := make(chan error, 1)

			err := srv.Start(errC)
			assert.Nil(t, err)
		})
	}

	srv.GracefulShutdown(ctx)
}
