package services

import (
	"checkout-case/domain"
	"checkout-case/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPromotionService_calculateSameSellerPromotion(t *testing.T) {

	tests := []struct {
		name     string
		cart     *domain.Cart
		expected *models.PromotionServiceResponse
		isError  bool
	}{
		{
			name: "calculateSameSellerPromotion failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1001,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1002,
						Quantity: 1,
						Price:    100,
					},
				},
				TotalPrice: 300,
			},
			expected: nil,
			isError:  true,
		},
		{
			name: "calculateSameSellerPromotion succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
				},
				TotalPrice: 300,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: sameSellerPromotion,
				TotalDiscount:      30,
				TotalPrice:         270,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := calculateSameSellerPromotion(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPromotionService_calculateCategoryPromotion(t *testing.T) {

	tests := []struct {
		name     string
		cart     *domain.Cart
		expected *models.PromotionServiceResponse
		isError  bool
	}{
		{
			name: "calculateCategoryPromotion failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    100,
					},
				},
				TotalPrice: 300,
			},
			expected: nil,
			isError:  true,
		},
		{
			name: "calculateCategoryPromotion succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1000,
						Quantity:   1,
						Price:      100,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      100,
					},
					{
						CategoryId: 1000,
						SellerId:   1002,
						Quantity:   1,
						Price:      100,
					},
					{
						CategoryId: 1000,
						SellerId:   1003,
						Quantity:   1,
						Price:      100,
					},
					{
						CategoryId: 1000,
						SellerId:   1004,
						Quantity:   1,
						Price:      100,
					},
				},
				TotalPrice: 500,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: categoryPromotion,
				TotalDiscount:      10,
				TotalPrice:         490,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := calculateCategoryPromotion(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPromotionService_calculateTotalPricePromotion(t *testing.T) {

	tests := []struct {
		name     string
		cart     *domain.Cart
		expected *models.PromotionServiceResponse
		isError  bool
	}{
		{
			name: "total price zero failure",
			cart: &domain.Cart{
				Items:      []*domain.Item{},
				TotalPrice: 0,
			},
			expected: nil,
			isError:  true,
		},
		{
			name: "more discount than total price",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      100,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      100,
					},
				},
				TotalPrice: 200,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      0,
				TotalPrice:         0,
			},
			isError: false,
		},
		{
			name: "total price less than 5k",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      1000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      1000,
					},
				},
				TotalPrice: 2000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      250,
				TotalPrice:         1750,
			},
			isError: false,
		},
		{
			name: "total price between 5k to 10k",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      3000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      3000,
					},
				},
				TotalPrice: 6000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      500,
				TotalPrice:         5500,
			},
			isError: false,
		},
		{
			name: "total price between 10k to 50k",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      13000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      3000,
					},
				},
				TotalPrice: 16000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      1000,
				TotalPrice:         15000,
			},
			isError: false,
		},
		{
			name: "total price greater than 50k",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      13000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1001,
						Quantity:   1,
						Price:      53000,
					},
				},
				TotalPrice: 66000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      2000,
				TotalPrice:         64000,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := calculateTotalPricePromotion(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPromotionService_FindBestPromotion(t *testing.T) {
	promo := NewPromotionService()

	tests := []struct {
		name     string
		cart     *domain.Cart
		expected *models.PromotionServiceResponse
		isError  bool
	}{
		{
			name: "find best promo sameSellerPromotion succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
					{
						CategoryId: categoryPromotionCategoryID,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
				},
				TotalPrice: 3000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: sameSellerPromotion,
				TotalDiscount:      300,
				TotalPrice:         2700,
			},
			isError: false,
		},
		{
			name: "find best promo sameSellerPromotion succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 20,
						Price:    100,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    1000,
					},
				},
				TotalPrice: 3000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: sameSellerPromotion,
				TotalDiscount:      300,
				TotalPrice:         2700,
			},
			isError: false,
		},
		{
			name: "find best promo sameSellerPromotion succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 10,
						Price:    1000,
					},
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    1000,
					},
				},
				TotalPrice: 11000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: sameSellerPromotion,
				TotalDiscount:      1100,
				TotalPrice:         9900,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := promo.FindBestPromotion(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
