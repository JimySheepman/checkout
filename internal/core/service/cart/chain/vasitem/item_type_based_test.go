package vasitem

import (
	"checkout-case/internal/core/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesItemTypeBased_IsAddVasItemValid(t *testing.T) {
	rule := &rulesItemTypeBased{}

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		isError    bool
	}{
		{
			name: "item type failure",
			item: &domain.Item{
				Type: domain.DigitalItem,
			},
			totalPrice: 0,
			isError:    true,
		},
		{
			name: "item type succeed",
			item: &domain.Item{
				Type: domain.DefaultItem,
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
