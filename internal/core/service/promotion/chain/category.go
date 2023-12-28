package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const (
	categoryPromotionID         = 5676
	categoryPromotionCategoryID = 3003
	categoryPromotionDiscount   = 0.05
)

type categoryPromotion struct {
}

func (p categoryPromotion) Promotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("calculate category promotion")

	var discount float64

	for _, item := range cart.Items {
		if item.CategoryId == categoryPromotionCategoryID {
			discount += item.Price * categoryPromotionDiscount * float64(item.Quantity)
		}
	}

	if discount != 0 {
		return &models.PromotionServiceResponse{
			AppliedPromotionID: categoryPromotionID,
			TotalDiscount:      discount,
			TotalPrice:         cart.TotalPrice - discount,
		}, nil
	}

	err := fmt.Errorf("category promotion cannot be applied")
	l.Error(err)

	return nil, err
}
