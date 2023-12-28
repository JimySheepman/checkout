package vasitem

import (
	"checkout-case/internal/core/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesSellerID_IsAddVasItemValid(t *testing.T) {
	rule := &rulesSellerID{}

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		isError    bool
	}{
		{
			name: "vasItem SellerId failure",
			item: &domain.Item{
				VasItems: []*domain.VasItem{
					{
						SellerId: 1,
					},
					{
						SellerId: vasItemSellerID,
					},
					{
						SellerId: vasItemSellerID,
					},
				},
			},
			totalPrice: 0,
			isError:    true,
		},
		{
			name: "vasItem SellerId succeed",
			item: &domain.Item{
				VasItems: []*domain.VasItem{
					{
						SellerId: vasItemSellerID,
					},
					{
						SellerId: vasItemSellerID,
					},
					{
						SellerId: vasItemSellerID,
					},
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
