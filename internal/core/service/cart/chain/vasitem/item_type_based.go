package vasitem

import (
	"checkout-case/internal/core/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

type rulesItemTypeBased struct {
}

func (p *rulesItemTypeBased) IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules item type based is start")

	if item.Type != domain.DefaultItem {
		return fmt.Errorf("item type error")
	}
	l.Info("rules item type based done")

	return nil
}
