package vasitem

import (
	"checkout-case/internal/core/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const maxTotalPrice = 500_000

type rulesTotalPrice struct {
}

func (p *rulesTotalPrice) IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error {
	return p.isTotalPriceValid(ctx, totalPrice+item.Price)
}

func (p *rulesTotalPrice) isTotalPriceValid(ctx context.Context, totalPrice float64) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules total price is start")

	if totalPrice > maxTotalPrice {
		return fmt.Errorf("total price invalid")
	}
	l.Info("rules total price done")

	return nil
}
