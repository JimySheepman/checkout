package chain

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
)

type Promoter interface {
	Promotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error)
}

var promoter = []Promoter{
	&sameSellerPromotion{},
	&categoryPromotion{},
	&totalPricePromotion{},
}

func GetPromoter() []Promoter {
	return promoter
}
