package services

import (
	"checkout-case/domain"
	"checkout-case/models"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
)

const (
	sameSellerPromotion         = 9909
	sameSellerPromotionDiscount = 0.9

	categoryPromotion           = 5676
	categoryPromotionCategoryID = 3003
	categoryPromotionDiscount   = 0.05

	totalPricePromotion             = 1232
	totalPricePromotionBaseDiscount = 250
	totalPricePromotionLimit5k      = 5000
	totalPricePromotionLimit10k     = 10000
	totalPricePromotionLimit50k     = 50000
)

type calculatePromotionFunction func(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error)

var calculatePromotions = []calculatePromotionFunction{
	calculateSameSellerPromotion,
	calculateCategoryPromotion,
	calculateTotalPricePromotion,
}

type promotionService struct {
}

func NewPromotionService() *promotionService {
	return &promotionService{}
}

func calculateSameSellerPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
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
			AppliedPromotionID: sameSellerPromotion,
			TotalDiscount:      cart.TotalPrice - newTotalPrice,
			TotalPrice:         newTotalPrice,
		}, nil
	}

	err := fmt.Errorf("same seller promotion cannot be applied")
	l.Error(err)

	return nil, err
}

func calculateCategoryPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
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
			AppliedPromotionID: categoryPromotion,
			TotalDiscount:      discount,
			TotalPrice:         cart.TotalPrice - discount,
		}, nil
	}

	err := fmt.Errorf("category promotion cannot be applied")
	l.Error(err)

	return nil, err
}

func calculateTotalPricePromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
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
				AppliedPromotionID: totalPricePromotion,
				TotalDiscount:      0,
				TotalPrice:         0,
			}, nil
		}

		return &models.PromotionServiceResponse{
			AppliedPromotionID: totalPricePromotion,
			TotalDiscount:      discount,
			TotalPrice:         t,
		}, nil
	}

	err := fmt.Errorf("total price promotion cannot be applied")
	l.Error(err)

	return nil, err
}

func (s *promotionService) FindBestPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("find best promotion")

	promotions := make([]*models.PromotionServiceResponse, 0, len(calculatePromotions))

	for _, calculatePromotion := range calculatePromotions {
		promotion, err := calculatePromotion(ctx, cart)
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
