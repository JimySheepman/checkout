package file

import (
	"checkout-case/internal/core/models"
	mock_handler "checkout-case/mocks"
	"checkout-case/pkg/customerr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fileHandlerMocks struct {
	mockCartService *mock_handler.MockCartService
}

func setupFileHandlerTest(t *testing.T) (context.Context, *fileHandlerMocks, *fileHandler) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &fileHandlerMocks{
		mockCartService: mock_handler.NewMockCartService(ctrl),
	}

	srv := NewFileHandler(mocks.mockCartService)

	return ctx, mocks, srv
}

func TestFileHandler_AddItemHandler(t *testing.T) {
	ctx, mocks, srv := setupFileHandlerTest(t)

	tests := []struct {
		name         string
		input        string
		expected     string
		isError      bool
		expectations func()
	}{
		{
			name:         "command Unmarshal failure",
			input:        "",
			expected:     "",
			isError:      true,
			expectations: func() {},
		},
		{
			name:     "AddItemToCart failure",
			input:    `{"command":"addItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: "",
			isError:  true,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddItemToCart(gomock.Any(), gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:     "AddItemHandler succeed",
			input:    `{"command":"addItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: "item was added to cart successfully",
			isError:  false,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddItemToCart(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.AddItemHandler(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileHandler_AddVasItemToItemHandler(t *testing.T) {
	ctx, mocks, srv := setupFileHandlerTest(t)

	tests := []struct {
		name         string
		input        string
		expected     string
		isError      bool
		expectations func()
	}{
		{
			name:         "command Unmarshal failure",
			input:        "",
			expected:     "",
			isError:      true,
			expectations: func() {},
		},
		{
			name:     "AddVasItemToItem failure",
			input:    `{"command":"addVasItemToItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: "",
			isError:  true,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddVasItemToItem(gomock.Any(), gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:     "AddVasItemToItemHandler succeed",
			input:    `{"command":"addVasItemToItem","payload":{"itemId":1,"categoryId":2,"sellerId":3,"price":4.0,"quantity":5}}`,
			expected: "vasItem was added to item successfully",
			isError:  false,
			expectations: func() {
				mocks.mockCartService.EXPECT().AddVasItemToItem(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.AddVasItemToItemHandler(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileHandler_RemoveItemHandler(t *testing.T) {
	ctx, mocks, srv := setupFileHandlerTest(t)

	tests := []struct {
		name         string
		input        string
		expected     string
		isError      bool
		expectations func()
	}{
		{
			name:         "command Unmarshal failure",
			input:        "",
			expected:     "",
			isError:      true,
			expectations: func() {},
		},
		{
			name:     "RemoveItemFromCart failure",
			input:    `{"command":"removeItem","payload":{"itemId":1}}`,
			expected: "",
			isError:  true,
			expectations: func() {
				mocks.mockCartService.EXPECT().RemoveItemFromCart(gomock.Any(), gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:     "RemoveItemHandler succeed",
			input:    `{"command":"removeItem","payload":{"itemId":1}}`,
			expected: "item was removed to cart successfully",
			isError:  false,
			expectations: func() {
				mocks.mockCartService.EXPECT().RemoveItemFromCart(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.RemoveItemHandler(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileHandler_ResetCartHandler(t *testing.T) {
	ctx, mocks, srv := setupFileHandlerTest(t)

	tests := []struct {
		name         string
		input        string
		expected     string
		isError      bool
		expectations func()
	}{
		{
			name:     "ResetCart failure",
			input:    `{"command":"resetCart"}`,
			expected: "",
			isError:  true,
			expectations: func() {
				mocks.mockCartService.EXPECT().ResetCart(gomock.Any()).Return(customerr.ErrTest)
			},
		},
		{
			name:     "ResetCartHandler succeed",
			input:    `{"command":"resetCart"}`,
			expected: "cart was reset successfully",
			isError:  false,
			expectations: func() {
				mocks.mockCartService.EXPECT().ResetCart(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.ResetCartHandler(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFileHandler_DisplayCartHandler(t *testing.T) {
	ctx, mocks, srv := setupFileHandlerTest(t)

	cart := &models.DisplayCartServiceResponse{}

	tests := []struct {
		name         string
		input        string
		expected     string
		isError      bool
		expectations func()
	}{
		{
			name:     "DisplayCart failure",
			input:    `{"command":"displayCart"}`,
			expected: "",
			isError:  true,
			expectations: func() {
				mocks.mockCartService.EXPECT().DisplayCart(gomock.Any()).Return(nil, customerr.ErrTest)
			},
		},
		{
			name:     "DisplayCartHandler succeed",
			input:    `{"command":"displayCart"}`,
			expected: cart.ToString(),
			isError:  false,
			expectations: func() {
				mocks.mockCartService.EXPECT().DisplayCart(gomock.Any()).Return(cart, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.DisplayCartHandler(ctx, tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
