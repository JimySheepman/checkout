package rules

import (
	"checkout-case/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesTotalPrice_IsAddVasItemValid(t *testing.T) {
	rule := &rulesTotalPrice{}

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		isError    bool
	}{
		{
			name: "max total price failure",
			item: &domain.Item{
				CategoryId: 12,
				Price:      11000,
			},
			totalPrice: 490000,
			isError:    true,
		},
		{
			name: "max total price succeed",
			item: &domain.Item{
				CategoryId: 12,
				Price:      11000,
			},
			totalPrice: 11000,
			isError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := rule.IsAddVasItemValid(context.TODO(), tt.item, tt.totalPrice)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestRulesTotalPrice_IsAddItemValid(t *testing.T) {
	rule := &rulesTotalPrice{}

	tests := []struct {
		name    string
		cart    *domain.Cart
		isError bool
	}{
		{
			name: "max total price failure",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    500001,
					},
				},
				TotalPrice: 500001,
			},
			isError: true,
		},
		{
			name: "max total price succeed",
			cart: &domain.Cart{
				Items: []*domain.Item{
					{
						SellerId: 1000,
						Quantity: 1,
						Price:    5001,
					},
				},
				TotalPrice: 5001,
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
