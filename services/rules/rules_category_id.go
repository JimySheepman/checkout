package rules

import (
	"checkout-case/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const (
	vasItemCategoryIDFurniture   = 1001
	vasItemCategoryIDElectronics = 3004
)

type rulesCategoryID struct {
}

func (p *rulesCategoryID) IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules item type based is start")

	if item.CategoryId == vasItemCategoryIDFurniture || item.CategoryId == vasItemCategoryIDElectronics {
		l.Info("rules item type based done")
		return nil
	}

	return fmt.Errorf("item category id error")
}
