package domain

import (
	"time"
)

const (
	DefaultItem = iota
	DigitalItem
)

type Item struct {
	ID         string     `db:"ID" bson:"_id" json:"id"`
	ItemId     int        `db:"ItemId" bson:"itemId" json:"itemId"`
	CategoryId int        `db:"CategoryId" bson:"categoryId" json:"categoryId"`
	SellerId   int        `db:"SellerId" bson:"sellerId" json:"sellerId"`
	CreatedAt  time.Time  `db:"CreatedAt" bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time  `db:"UpdatedAt" bson:"updatedAt" json:"updatedAt"`
	Type       int        `db:"Type" bson:"type" json:"type"`
	Price      float64    `db:"Price" bson:"price" json:"price"`
	Quantity   int        `db:"Quantity" bson:"quantity" json:"quantity"`
	VasItems   []*VasItem `db:"VasItems" bson:"vasItems" json:"vasItems"`
}

type VasItem struct {
	ID         string    `db:"ID" bson:"_id" json:"id"`
	VasItemId  int       `db:"VasItemId" bson:"vasItemId" json:"vasItemId"`
	CategoryId int       `db:"CategoryId" bson:"categoryId" json:"categoryId"`
	SellerId   int       `db:"SellerId" bson:"sellerId" json:"sellerId"`
	CreatedAt  time.Time `db:"CreatedAt" bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time `db:"UpdatedAt" bson:"updatedAt" json:"updatedAt"`
	Price      float64   `db:"Price" bson:"price" json:"price"`
	Quantity   int       `db:"Quantity" bson:"quantity" json:"quantity"`
}
