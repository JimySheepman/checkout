package promotion

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
						CategoryId: 3003,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
					{
						CategoryId: 3003,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
					{
						CategoryId: 3003,
						SellerId:   1000,
						Quantity:   1,
						Price:      1000,
					},
				},
				TotalPrice: 3000,
			},
			expected: &models.PromotionServiceResponse{
				AppliedPromotionID: 9909,
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
				AppliedPromotionID: 9909,
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
				AppliedPromotionID: 9909,
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
