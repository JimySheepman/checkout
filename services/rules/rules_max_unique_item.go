package rules

import (
	"checkout-case/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const maxUniqueItemCount = 10

type rulesMaxUniqueItem struct {
}

func (p *rulesMaxUniqueItem) IsAddItemValid(ctx context.Context, cart *domain.Cart) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules max unique item count is start")

	if len(cart.Items) > maxUniqueItemCount {
		return fmt.Errorf("unieque item count invalid")
	}
	l.Info("rules max unique item count done")

	return nil
}
