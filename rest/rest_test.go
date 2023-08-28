//go:generate mockgen -destination=./mocks/rest_mock.go -source=rest.go checkout-case/rest

package rest

import (
	"checkout-case/config"
	mock_rest "checkout-case/rest/mocks"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type restServerMocks struct {
	mockRestHandlerClient *mock_rest.MockrestHandlerClient
}

func setupRestServerTest(t *testing.T) (context.Context, *restServerMocks, *RestServer) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &restServerMocks{
		mockRestHandlerClient: mock_rest.NewMockrestHandlerClient(ctrl),
	}

	srv := NewRestServer(mocks.mockRestHandlerClient)

	return ctx, mocks, srv
}

func TestRestServer_Start(t *testing.T) {
	ctx, mocks, srv := setupRestServerTest(t)

	cfg, err := config.Load()
	require.Nil(t, err)

	cfg.Server.RestServer.PprofEnable = 1
	cfg.Server.RestServer.Addr = ":9090"

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
