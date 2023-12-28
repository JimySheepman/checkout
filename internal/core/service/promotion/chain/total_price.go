package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const (
	totalPricePromotionID           = 1232
	totalPricePromotionBaseDiscount = 250
	totalPricePromotionLimit5k      = 5000
	totalPricePromotionLimit10k     = 10000
	totalPricePromotionLimit50k     = 50000
)

type totalPricePromotion struct{}

func (p totalPricePromotion) Promotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("calculate total price promotion")

	var (
		discount float64
		tp       = cart.TotalPrice
	)

	switch {
	case tp > 0 && totalPricePromotionLimit5k > tp:
		discount = totalPricePromotionBaseDiscount
	case tp >= totalPricePromotionLimit5k && totalPricePromotionLimit10k > tp:
		discount = totalPricePromotionBaseDiscount * 2
	case tp >= totalPricePromotionLimit10k && totalPricePromotionLimit50k > tp:
		discount = totalPricePromotionBaseDiscount * 4
	case tp >= totalPricePromotionLimit50k:
		discount = totalPricePromotionBaseDiscount * 8
	default:
		discount = 0
	}

	if discount != 0 {
		t := cart.TotalPrice - discount
		if 0 > t {
			return &models.PromotionServiceResponse{
				AppliedPromotionID: totalPricePromotionID,
				TotalDiscount:      0,
				TotalPrice:         0,
			}, nil
		}

		return &models.PromotionServiceResponse{
			AppliedPromotionID: totalPricePromotionID,
			TotalDiscount:      discount,
			TotalPrice:         t,
		}, nil
	}

	err := fmt.Errorf("total price promotion cannot be applied")
	l.Error(err)

	return nil, err
}
