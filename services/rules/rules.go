package rules

import (
	"checkout-case/domain"
	"context"
)

type Ruler interface {
	IsAddItemValid(ctx context.Context, cart *domain.Cart) error
}

var ruler = []Ruler{
	&rulesTotalPrice{},            // total price <= 500k
	&rulesAllItemCount{},          // total products <= 30
	&rulesMaxUniqueItem{},         // max unique item <= 10
	&rulesItemTypeBasedQuantity{}, // digital item quantity type controller
}

func GetRulers() []Ruler {
	return ruler
}

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
