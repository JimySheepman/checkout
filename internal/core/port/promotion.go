//go:generate mockgen -destination=../../../mocks/promotion_mock.go -source=promotion.go

package port

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
)

type PromotionServiceClient interface {
	FindBestPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error)
}
