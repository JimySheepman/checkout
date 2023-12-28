package promotion

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"checkout-case/internal/core/service/promotion/chain"
	"checkout-case/pkg/logger"
	"context"
)

type promotionService struct {
}

func NewPromotionService() *promotionService {
	return &promotionService{}
}

func (s *promotionService) FindBestPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("find best promotion")

	promotions := make([]*models.PromotionServiceResponse, 0, len(chain.GetPromoter()))

	for _, promoter := range chain.GetPromoter() {
		promotion, err := promoter.Promotion(ctx, cart)
		if err != nil {
			continue
		}

		promotions = append(promotions, promotion)
	}

	var (
		chooseBestDiscount       float64
		chooseBestPromotionIndex int
	)

	for i, promotion := range promotions {
		if promotion.TotalDiscount > chooseBestDiscount {
			chooseBestDiscount = promotion.TotalDiscount
			chooseBestPromotionIndex = i
		}
	}

	return promotions[chooseBestPromotionIndex], nil
}
