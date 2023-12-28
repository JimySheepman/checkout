package item

import (
	"checkout-case/internal/core/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesMaxUniqueItem_IsAddItemValid(t *testing.T) {
	rule := &rulesMaxUniqueItem{}

	tests := []struct {
		name    string
		cart    *domain.Cart
		isError bool
	}{
		{
			name: "max unique item count failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						ItemId: 1,
					},
					{
						ItemId: 2,
					},
					{
						ItemId: 3,
					},
					{
						ItemId: 4,
					},
					{
						ItemId: 5,
					},
					{
						ItemId: 6,
					},
					{
						ItemId: 7,
					},
					{
						ItemId: 8,
					}, {
						ItemId: 9,
					},
					{
						ItemId: 10,
					},
					{
						ItemId: 11,
					},
					{
						ItemId: 12,
					},
				},
			},
			isError: true,
		},
		{
			name: "IsAddItemValid succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						ItemId:   1,
						SellerId: 1000,
						Quantity: 1,
						Price:    10,
					},
					{
						ItemId:   2,
						SellerId: 1001,
						Quantity: 3,
						Price:    10,
					},
				},
				TotalPrice: 40,
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := rule.IsAddItemValid(context.TODO(), tt.cart)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
