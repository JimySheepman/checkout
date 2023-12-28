package http

import (
	"bytes"
	mock_handler "checkout-case/handler/mocks"
	"checkout-case/models"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRestHandlerTest(t *testing.T) (*echo.Echo, *commonMocks, *restHandler) {
	ctrl, _ := gomock.WithContext(context.Background(), t)

	mocks := &commonMocks{
		mockCartService: mock_handler.NewMockcartService(ctrl),
	}

	srv := NewRestHandler(mocks.mockCartService)
	e := echo.New()

	return e, mocks, srv
}

func TestRestHandler_HealthCheckHandler(t *testing.T) {
	e, _, srv := setupRestHandlerTest(t)

	tests := []struct {
		name string
	}{
		{
			name: "HealthCheckHandler succeed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := srv.HealthCheckHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestRestHandler_AddItemHandler(t *testing.T) {
	e, mocks, srv := setupRestHandlerTest(t)

	item := &models.AddItemRestRequest{}
	testItem, err := json.Marshal(item)
	require.Nil(t, err)

	tests := []struct {
		name         string
		body         []byte
		expectations func()
	}{
		{
			name: "json.NewDecoder failure",
			body: []byte("test"),
			expectations: func() {
			},
		},
		{
			name: "AddItemToCart failure",
			body: testItem,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddItemToCart(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
		{
			name: "AddItemHandler succeed",
			body: testItem,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddItemToCart(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := srv.AddItemHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestRestHandler_AddVasItemToItemHandler(t *testing.T) {
	e, mocks, srv := setupRestHandlerTest(t)

	item := &models.AddVasItemToItemRestRequest{}
	testItem, err := json.Marshal(item)
	require.Nil(t, err)

	tests := []struct {
		name         string
		body         []byte
		expectations func()
	}{
		{
			name: "json.NewDecoder failure",
			body: []byte("test"),
			expectations: func() {
			},
		},
		{
			name: "AddVasItemToItem failure",
			body: testItem,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddVasItemToItem(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
		{
			name: "AddVasItemToItemHandler succeed",
			body: testItem,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddVasItemToItem(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := srv.AddVasItemToItemHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestRestHandler_RemoveItemHandler(t *testing.T) {
	e, mocks, srv := setupRestHandlerTest(t)

	tests := []struct {
		name         string
		param        string
		body         []byte
		expectations func()
	}{
		{
			name:  "strconv.Atoi failure",
			param: "test",
			expectations: func() {
			},
		},
		{
			name:  "RemoveItemFromCart failure",
			param: "12",
			expectations: func() {
				mocks.mockCartService.EXPECT().RemoveItemFromCart(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
		{
			name:  "RemoveItemHandler succeed",
			param: "12",
			expectations: func() {
				mocks.mockCartService.EXPECT().RemoveItemFromCart(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("itemId")
			ctx.SetParamValues(tt.param)

			err := srv.RemoveItemHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestRestHandler_ResetCartHandler(t *testing.T) {
	e, mocks, srv := setupRestHandlerTest(t)

	tests := []struct {
		name         string
		body         []byte
		expectations func()
	}{
		{
			name: "ResetCart failure",
			body: nil,
			expectations: func() {
				mocks.mockCartService.EXPECT().ResetCart(gomock.Any()).Return(ErrTest)
			},
		},
		{
			name: "ResetCartHandler succeed",
			body: nil,
			expectations: func() {
				mocks.mockCartService.EXPECT().ResetCart(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := srv.ResetCartHandler(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestRestHandler_DisplayCartHandler(t *testing.T) {
	e, mocks, srv := setupRestHandlerTest(t)

	tests := []struct {
		name         string
		body         []byte
		expectations func()
	}{
		{
			name: "DisplayCart failure",
			body: nil,
			expectations: func() {
				mocks.mockCartService.EXPECT().DisplayCart(gomock.Any()).Return(nil, ErrTest)
			},
		},
		{
			name: "DisplayCartHandler succeed",
			body: nil,
			expectations: func() {
				mocks.mockCartService.EXPECT().DisplayCart(gomock.Any()).Return(&models.DisplayCartServiceResponse{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(tt.body))
			req.Header.Set(echo.HeaderXRequestID, "test-request-id")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			err := srv.DisplayCartHandler(ctx)
			assert.Nil(t, err)
		})
	}
}
