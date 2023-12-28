package item

import (
	"checkout-case/internal/core/domain"
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
