package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cart struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
	Items      []*Item            `bson:"items"`
	TotalPrice float64            `bson:"totalPrice"`
}
