package vasitem

import (
	"checkout-case/internal/core/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const maxVasItemCount = 3

type rulesVasItemCount struct {
}

func (p *rulesVasItemCount) IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules vas item count is start")
	tmpQuantity := 0

	for _, vasItem := range item.VasItems {
		tmpQuantity += vasItem.Quantity
	}

	if tmpQuantity > item.Quantity*maxVasItemCount {
		return fmt.Errorf("vasItem len error")
	}

	if len(item.VasItems) > maxVasItemCount {
		return fmt.Errorf("vasItem len error")
	}
	l.Info("rules vas item count done")

	return nil
}
