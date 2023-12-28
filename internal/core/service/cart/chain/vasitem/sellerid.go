package vasitem

import (
	"checkout-case/internal/core/domain"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const vasItemSellerID = 5003

type rulesSellerID struct {
}

func (p *rulesSellerID) IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("rules sellerID is start")

	for _, vasItem := range item.VasItems {
		if vasItem.SellerId != vasItemSellerID {
			return fmt.Errorf("seller id error")
		}
	}
	l.Info("rules sellerID done")

	return nil
}
