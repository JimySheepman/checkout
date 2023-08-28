package rules

import (
	"checkout-case/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const maxDigitalItemCount = 5

type rulesItemTypeBasedQuantity struct {
}

func (p *rulesItemTypeBasedQuantity) IsAddItemValid(ctx context.Context, cart *domain.Cart) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules item type based quantity is start")

	var count int

	for _, item := range cart.Items {
		if item.Type == domain.DigitalItem {
			count += item.Quantity
		}
	}

	if count > maxDigitalItemCount {
		return fmt.Errorf("digital item is not added")
	}
	l.Info("rules item type based quantity done")

	return nil
}
