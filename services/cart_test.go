//go:generate mockgen -destination=./mocks/cart_mock.go -source=cart.go checkout-case/services

package services

import (
	"checkout-case/domain"
	"checkout-case/models"
	mock_services "checkout-case/services/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ErrTest = errors.New("test error")

type cartServiceMocks struct {
	mockCartRepository         *mock_services.MockcartRepository
	mockPromotionServiceClient *mock_services.MockpromotionServiceClient
}

func setupCartServiceTest(t *testing.T) (context.Context, *cartServiceMocks, *cartService) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	mocks := &cartServiceMocks{
		mockCartRepository:         mock_services.NewMockcartRepository(ctrl),
		mockPromotionServiceClient: mock_services.NewMockpromotionServiceClient(ctrl),
	}

	srv := NewCartService(mocks.mockCartRepository, mocks.mockPromotionServiceClient)

	return ctx, mocks, srv
}

func TestCartService_AddItemToCart(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		req          *models.AddItemServiceRequest
		isError      bool
		expectations func()
	}{
		{
			name: "s.addItemToCartValidation failure",
			req: &models.AddItemServiceRequest{
				CategoryID: vasItemCategoryID,
			},
			isError:      true,
			expectations: func() {},
		},
		{
			name: "cartRepository.GetCart failure",
			req: &models.AddItemServiceRequest{
				CategoryID: 1,
				SellerID:   1,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(nil, ErrTest)
			},
		},
		{
			name: "s.addItemToCartRuleChain failure",
			req: &models.AddItemServiceRequest{
				CategoryID: 1,
				SellerID:   1,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							CategoryId: 1001,
							Type:       domain.DigitalItem,
							Price:      100,
							Quantity:   6,
						},
					},
					TotalPrice: 150,
				}, nil)
			},
		},
		{
			name: "s.sameItemProcess failure",
			req: &models.AddItemServiceRequest{
				CategoryID: 1,
				SellerID:   1,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							CategoryId: 1001,
							Type:       domain.DigitalItem,
							Price:      100,
							Quantity:   3,
						},
					},
					TotalPrice: 150,
				}, nil)
				mocks.mockCartRepository.EXPECT().UpdateItemQuantity(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
		{
			name: "cartRepository.AddItem failure",
			req: &models.AddItemServiceRequest{
				ItemID:     1,
				CategoryID: 1,
				SellerID:   1,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							ItemId:     2,
							CategoryId: 1001,
							Type:       domain.DigitalItem,
							Price:      100,
							Quantity:   3,
						},
					},
					TotalPrice: 150,
				}, nil)
				mocks.mockCartRepository.EXPECT().UpdateItemQuantity(gomock.Any(), gomock.Any()).Return(nil)
				mocks.mockCartRepository.EXPECT().AddItem(gomock.Any()).Return(ErrTest)
			},
		},
		{
			name: "AddItemToCart succeed",
			req: &models.AddItemServiceRequest{
				ItemID:     1,
				CategoryID: 1,
				SellerID:   1,
			},
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							ItemId:     2,
							CategoryId: 1001,
							Type:       domain.DigitalItem,
							Price:      100,
							Quantity:   3,
						},
					},
					TotalPrice: 150,
				}, nil)
				mocks.mockCartRepository.EXPECT().UpdateItemQuantity(gomock.Any(), gomock.Any()).Return(nil)
				mocks.mockCartRepository.EXPECT().AddItem(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.AddItemToCart(ctx, tt.req)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_addItemToCartValidation(t *testing.T) {
	_, _, srv := setupCartServiceTest(t)

	tests := []struct {
		name    string
		req     *models.AddItemServiceRequest
		isError bool
	}{
		{
			name: "req.CategoryID failure",
			req: &models.AddItemServiceRequest{
				CategoryID: vasItemCategoryID,
				SellerID:   1,
			},
			isError: true,
		},
		{
			name: "req.SellerID failure",
			req: &models.AddItemServiceRequest{
				CategoryID: 1,
				SellerID:   vasItemSellerID,
			},
			isError: true,
		},
		{
			name: "addItemToCartValidation succeed",
			req: &models.AddItemServiceRequest{
				CategoryID: 1,
				SellerID:   1,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := srv.addItemToCartValidation(tt.req)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_sameItemProcess(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		req          *models.AddItemServiceRequest
		item         *domain.Item
		ok           bool
		isError      bool
		expectations func()
	}{
		{
			name:    "cartRepository.UpdateItemQuantity failure",
			req:     &models.AddItemServiceRequest{},
			item:    &domain.Item{},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().UpdateItemQuantity(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
		{
			name:    "sameItemProcess succeed",
			req:     &models.AddItemServiceRequest{},
			item:    &domain.Item{},
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().UpdateItemQuantity(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.sameItemProcess(ctx, tt.item, tt.req)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_addItemToCartRuleChain(t *testing.T) {
	ctx, _, srv := setupCartServiceTest(t)

	tests := []struct {
		name       string
		cart       *domain.Cart
		totalPrice float64
		expected   bool
	}{
		{
			name: "rule.IsAddItemValid failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: 1001,
						Type:       domain.DigitalItem,
						Price:      100,
						Quantity:   6,
					},
				},
				TotalPrice: 150,
			},
			expected: false,
		},
		{
			name: "addItemToCartRuleChain succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: 1001,
						Type:       domain.DefaultItem,
						Price:      100,
						Quantity:   1,
						VasItems: []*domain.VasItem{
							{
								SellerId: 5003,
								Price:    50,
								Quantity: 1,
							},
						},
					},
				},
				TotalPrice: 150,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := srv.addItemToCartRuleChain(ctx, tt.cart)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCartService_AddVasItemToItem(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		req          *models.AddVasItemToItemServiceRequest
		isError      bool
		expectations func()
	}{
		{
			name:    "cartRepository.FindItemByItemIdFromCart failure",
			req:     &models.AddVasItemToItemServiceRequest{},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(nil, ErrTest)
			},
		},
		{
			name:    "s.addVasItemToItemValidation failure",
			req:     &models.AddVasItemToItemServiceRequest{},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{}, nil)
			},
		},
		{
			name: "cartRepository.GetCart failure",
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
				Price:      20,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{
					Type:  domain.DefaultItem,
					Price: 120,
				}, nil)
				mocks.mockCartRepository.EXPECT().GetCart().Return(nil, ErrTest)
			},
		},
		{
			name: "s.addVasItemToItemRuleChain failure",
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
				Price:      20,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{
					Type:  domain.DefaultItem,
					Price: 120,
				}, nil)
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							CategoryId: 1001,
							Type:       domain.DefaultItem,
							Price:      100,
							Quantity:   1,
							VasItems: []*domain.VasItem{
								{
									SellerId: 5003,
									Price:    50,
									Quantity: 1,
								},
							},
						},
					},
					TotalPrice: 150,
				}, nil)
			},
		},
		{
			name: "cartRepository.AddVasItemToItemByItemID failure",
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
			},
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{
					CategoryId: 1001,
					Type:       domain.DefaultItem,
					Quantity:   1,
					Price:      150,
				}, nil)
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							CategoryId: 1001,
							Type:       domain.DigitalItem,
							Price:      100,
							Quantity:   1,
						},
					},
					TotalPrice: 100,
				}, nil)
				mocks.mockCartRepository.EXPECT().AddVasItemToItemByItemID(gomock.Any(), gomock.Any()).Return(ErrTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.AddVasItemToItem(ctx, tt.req)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_addVasItemToItemValidation(t *testing.T) {
	_, _, srv := setupCartServiceTest(t)

	tests := []struct {
		name    string
		item    *domain.Item
		req     *models.AddVasItemToItemServiceRequest
		isError bool
	}{
		{
			name: "item.Price failure",
			item: &domain.Item{
				Type:  domain.DigitalItem,
				Price: 12,
			},
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
				Price:      15,
			},
			isError: true,
		},
		{
			name: "item.Type failure",
			item: &domain.Item{
				Type:  domain.DigitalItem,
				Price: 120,
			},
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
				Price:      15,
			},
			isError: true,
		},
		{
			name: "req.CategoryID failure",
			item: &domain.Item{
				Type:  domain.DefaultItem,
				Price: 120,
			},
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: 12,
				Price:      15,
			},
			isError: true,
		},
		{
			name: "addVasItemToItemValidation succeed",
			item: &domain.Item{
				Type:  domain.DefaultItem,
				Price: 120,
			},
			req: &models.AddVasItemToItemServiceRequest{
				CategoryID: vasItemCategoryID,
				Price:      15,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := srv.addVasItemToItemValidation(tt.item, tt.req)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_addVasItemToItemRuleChain(t *testing.T) {
	ctx, _, srv := setupCartServiceTest(t)

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		expected   bool
	}{
		{
			name: "rule.IsAddVasItemValid failure",
			item: &domain.Item{
				CategoryId: 1001,
				Type:       domain.DigitalItem,
				Price:      100,
				Quantity:   1,
				VasItems: []*domain.VasItem{
					{
						SellerId: 5003,
						Price:    50,
						Quantity: 1,
					},
				},
			},
			totalPrice: 150,
			expected:   false,
		},
		{
			name: "addVasItemToItemRuleChain succeed",
			item: &domain.Item{
				CategoryId: 1001,
				Type:       domain.DefaultItem,
				Price:      100,
				Quantity:   1,
				VasItems: []*domain.VasItem{
					{
						SellerId: 5003,
						Price:    50,
						Quantity: 1,
					},
				},
			},
			totalPrice: 150,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := srv.addVasItemToItemRuleChain(ctx, tt.item, tt.totalPrice)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCartService_RemoveItemFromCart(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		itemId       int
		isError      bool
		expectations func()
	}{
		{
			name:    "cartRepository.FindItemByItemIdFromCart failure",
			itemId:  1,
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(nil, ErrTest)
			},
		},
		{
			name:    "cartRepository.RemoveItem failure",
			itemId:  1,
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{}, nil)
				mocks.mockCartRepository.EXPECT().RemoveItem(gomock.Any()).Return(ErrTest)
			},
		},
		{
			name:    "cartRepository.RemoveItem failure",
			itemId:  1,
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().FindItemByItemIdFromCart(gomock.Any()).Return(&domain.Item{}, nil)
				mocks.mockCartRepository.EXPECT().RemoveItem(gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.RemoveItemFromCart(ctx, tt.itemId)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_ResetCart(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		isError      bool
		expectations func()
	}{
		{
			name:    "cartRepository.ResetCart failure",
			isError: true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().ResetCart().Return(ErrTest)
			},
		},
		{
			name:    "ResetCart succeed",
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().ResetCart().Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			err := srv.ResetCart(ctx)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCartService_DisplayCart(t *testing.T) {
	ctx, mocks, srv := setupCartServiceTest(t)

	tests := []struct {
		name         string
		expected     *models.DisplayCartServiceResponse
		isError      bool
		expectations func()
	}{
		{
			name:     "cartRepository.GetCart failure",
			expected: nil,
			isError:  true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(nil, ErrTest)
			},
		},
		{
			name: "TotalPrice is zero succeed",
			expected: &models.DisplayCartServiceResponse{
				Items:              nil,
				TotalPrice:         0,
				AppliedPromotionID: 0,
				TotalDiscount:      0,
			},
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{}, nil)
			},
		},
		{
			name:     "FindBestPromotion failure",
			expected: nil,
			isError:  true,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							Price:    10,
							Quantity: 1,
						},
					},
					TotalPrice: 10,
				}, nil)
				mocks.mockPromotionServiceClient.EXPECT().FindBestPromotion(gomock.Any(), gomock.Any()).Return(nil, ErrTest)
			},
		},
		{
			name: "FindBestPromotion failure",
			expected: &models.DisplayCartServiceResponse{
				Items: []models.ItemServiceResponse{
					{
						Price:    10,
						Quantity: 1,
					},
				},
				TotalPrice:         9,
				AppliedPromotionID: 1,
				TotalDiscount:      1,
			},
			isError: false,
			expectations: func() {
				mocks.mockCartRepository.EXPECT().GetCart().Return(&domain.Cart{
					Items: []*domain.Item{
						{
							Price:    10,
							Quantity: 1,
						},
					},
					TotalPrice: 10,
				}, nil)
				mocks.mockPromotionServiceClient.EXPECT().FindBestPromotion(gomock.Any(), gomock.Any()).Return(&models.PromotionServiceResponse{
					AppliedPromotionID: 1,
					TotalDiscount:      1,
					TotalPrice:         9,
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := srv.DisplayCart(ctx)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
