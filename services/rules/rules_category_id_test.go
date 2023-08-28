package rules

import (
	"checkout-case/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesCategoryID_IsAddVasItemValid(t *testing.T) {
	rule := &rulesCategoryID{}

	tests := []struct {
		name       string
		item       *domain.Item
		totalPrice float64
		isError    bool
	}{
		{
			name: "categoryId unknown failure",
			item: &domain.Item{
				CategoryId: 12,
			},
			totalPrice: 0,
			isError:    true,
		},
		{
			name: "categoryId vasItemCategoryIDFurniture succeed",
			item: &domain.Item{
				CategoryId: vasItemCategoryIDFurniture,
			},
			totalPrice: 0,
			isError:    false,
		},
		{
			name: "categoryId vasItemCategoryIDElectronics succeed",
			item: &domain.Item{
				CategoryId: vasItemCategoryIDElectronics,
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
