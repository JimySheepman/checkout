package vasitem

import (
	"checkout-case/internal/core/domain"
	"context"
)

type VasRuler interface {
	IsAddVasItemValid(ctx context.Context, item *domain.Item, totalPrice float64) error
}

var vasRuler = []VasRuler{
	&rulesTotalPrice{},    // total price <= 500k
	&rulesVasItemCount{},  // vasItem count <=3
	&rulesItemTypeBased{}, // item type check
	&rulesSellerID{},      // sellerID check
	&rulesCategoryID{},    // CategoryID check
}

func GetVasRulers() []VasRuler {
	return vasRuler
}
