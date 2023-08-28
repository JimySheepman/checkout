package rules

import (
	"checkout-case/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesVasItemCount_IsAddVasItemValid(t *testing.T) {
	rule := &rulesVasItemCount{}

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		isError    bool
	}{
		{
			name: "max vasItem quantity count failure",
			item: &domain.Item{
				Quantity: 3,
				VasItems: []*domain.VasItem{
					{
						Quantity: 10,
					},
				},
			},
			totalPrice: 0,
			isError:    true,
		},
		{
			name: "max vasItem count failure",
			item: &domain.Item{
				VasItems: []*domain.VasItem{
					{},
					{},
					{},
					{},
				},
			},
			totalPrice: 0,
			isError:    true,
		},
		{
			name: "max vasItem count succeed",
			item: &domain.Item{
				VasItems: []*domain.VasItem{
					{},
					{},
					{},
				},
			},
			totalPrice: 0,
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
