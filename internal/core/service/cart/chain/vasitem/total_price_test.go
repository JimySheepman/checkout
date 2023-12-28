package vasitem

import (
	"checkout-case/internal/core/domain"
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
