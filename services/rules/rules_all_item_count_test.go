package rules

import (
	"checkout-case/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesAllItemCount_IsAddItemValid(t *testing.T) {
	rule := &rulesAllItemCount{}

	tests := []struct {
		name    string
		cart    *domain.Cart
		isError bool
	}{
		{
			name: "max item count failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 10,
						Price:    1,
					},
					{
						SellerId: 1001,
						Quantity: 30,
						Price:    1,
					},
				},
				TotalPrice: 40,
			},
			isError: true,
		},
		{
			name: "IsAddItemValid succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    10,
					},
					{
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
