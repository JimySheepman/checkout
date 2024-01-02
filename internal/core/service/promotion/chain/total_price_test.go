package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPromotionService_calculateTotalPricePromotion(t *testing.T) {
	prom := &totalPricePromotion{}

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
				AppliedPromotionID: totalPricePromotionID,
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
				AppliedPromotionID: totalPricePromotionID,
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
				AppliedPromotionID: totalPricePromotionID,
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
				AppliedPromotionID: totalPricePromotionID,
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
				AppliedPromotionID: totalPricePromotionID,
				TotalDiscount:      2000,
				TotalPrice:         64000,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := prom.Promotion(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
