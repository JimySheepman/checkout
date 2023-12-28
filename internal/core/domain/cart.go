package domain

import (
	"time"
)

type Cart struct {
	ID         string    `db:"ID" bson:"_id" json:"id"`
	CreatedAt  time.Time `db:"CreatedAt" bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time `db:"UpdatedAt" bson:"updatedAt"  json:"updatedAt"`
	Items      []*Item   `db:"Items" bson:"items"  json:"items"`
	TotalPrice float64   `db:"TotalPrice" bson:"totalPrice" json:"TotalPrice"`
}
