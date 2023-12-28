package item

import (
	"checkout-case/internal/core/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const maxItemCount = 30

type rulesAllItemCount struct {
}

func (p *rulesAllItemCount) IsAddItemValid(ctx context.Context, cart *domain.Cart) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules all item count is start")

	var count int

	for _, item := range cart.Items {
		count += item.Quantity
	}

	if count > maxItemCount {
		return fmt.Errorf("item count invalid")
	}
	l.Info("rules all item count done")

	return nil
}
