package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const (
	sameSellerPromotionID       = 9909
	sameSellerPromotionDiscount = 0.9
)

type sameSellerPromotion struct{}

func (p sameSellerPromotion) Promotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("calculate same seller promotion")

	sellerIDs := make(map[int]bool)

	for _, item := range cart.Items {
		if _, value := sellerIDs[item.SellerId]; !value {
			sellerIDs[item.SellerId] = true
		}
	}

	if len(sellerIDs) == 1 {
		newTotalPrice := cart.TotalPrice * sameSellerPromotionDiscount

		return &models.PromotionServiceResponse{
			AppliedPromotionID: sameSellerPromotionID,
			TotalDiscount:      cart.TotalPrice - newTotalPrice,
			TotalPrice:         newTotalPrice,
		}, nil
	}

	err := fmt.Errorf("same seller promotion cannot be applied")
	l.Error(err)

	return nil, err
}
