package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	DefaultItem = iota
	DigitalItem
)

type Item struct {
	ID         primitive.ObjectID `bson:"_id"`
	ItemId     int                `bson:"itemId"`
	CategoryId int                `bson:"categoryId"`
	SellerId   int                `bson:"sellerId"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	Type       int                `bson:"type"`
	Price      float64            `bson:"price"`
	Quantity   int                `bson:"quantity"`
	VasItems   []*VasItem         `bson:"vasItems"`
}

type VasItem struct {
	ID         primitive.ObjectID `bson:"_id"`
	VasItemId  int                `bson:"vasItemId"`
	CategoryId int                `bson:"categoryId"`
	SellerId   int                `bson:"sellerId"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	Price      float64            `bson:"price"`
	Quantity   int                `bson:"quantity"`
}
