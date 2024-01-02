package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSameSellerPromotion_Promotion(t *testing.T) {
	prom := &sameSellerPromotion{}

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
				AppliedPromotionID: sameSellerPromotionID,
				TotalDiscount:      30,
				TotalPrice:         270,
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
