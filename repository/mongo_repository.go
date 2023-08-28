package repository

import (
	"checkout-case/config"
	"checkout-case/domain"
	"checkout-case/models"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var cartID primitive.ObjectID

type cartRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewCartRepository() *cartRepository {
	l := logger.GetLogger().Sugar()

	cfg := config.Config.MongoDB

	// example: uri := "mongodb://root:example@localhost:27017/?timeoutMS=5000"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?timeoutMS=%d", cfg.User, cfg.Password, cfg.Addr, cfg.Port, cfg.Timeout)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	l.Info("mongodb successfully connected")

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	l.Info("mongodb successfully pinged")

	return &cartRepository{
		client:     client,
		collection: client.Database(cfg.Name).Collection(cfg.Collection),
	}
}

func (r *cartRepository) Create() error {
	l := logger.GetLogger().Sugar()
	now := time.Now()

	defer func() {
		l.Infow("successfully created the cart", "duration", time.Since(now))
	}()

	c := &domain.Cart{
		ID:         primitive.NewObjectID(),
		CreatedAt:  now,
		UpdatedAt:  now,
		Items:      make([]*domain.Item, 0),
		TotalPrice: 0,
	}

	if _, err := r.collection.InsertOne(context.TODO(), c); err != nil {
		return err
	}
	l.Info("initial cart document has been created with the id : ", c.ID.Hex())

	cartID = c.ID

	return nil
}

func (r *cartRepository) GetCart() (*domain.Cart, error) {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully found the cart", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	cart := &domain.Cart{}
	if err := r.collection.FindOne(context.Background(), filter).Decode(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *cartRepository) AddItem(item *domain.Item) error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully added item to the cart", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$push",
			Value: bson.M{
				"items": item,
			},
		},
		{
			Key: "$set",
			Value: bson.M{
				"updatedAt": time.Now(),
			},
		},
		{
			Key: "$inc",
			Value: bson.M{
				"totalPrice": float64(item.Quantity) * item.Price,
			},
		},
	}

	if _, err := r.collection.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}

	return nil
}

func (r *cartRepository) UpdateItemQuantity(item *domain.Item, req *models.AddItemServiceRequest) error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully updated item to the cart", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"items.$[selectedItem].quantity":  item.Quantity,
				"items.$[selectedItem].updatedAt": now,
				"updatedAt":                       now,
			},
		},
		{
			Key: "$inc",
			Value: bson.M{
				"totalPrice": float64(req.Quantity) * item.Price,
			},
		},
	}

	opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{
				{
					Key:   "selectedItem._id",
					Value: item.ID,
				},
			},
		},
	})

	res := r.collection.FindOneAndUpdate(context.Background(), filter, update, opts)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (r *cartRepository) AddVasItemToItemByItemID(itemId primitive.ObjectID, vasItem *domain.VasItem) error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully added vasItem to the item", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$push",
			Value: bson.M{
				"items.$[selectedItem].vasItems": vasItem,
			},
		},
		{
			Key: "$set",
			Value: bson.M{
				"updatedAt":                       now,
				"items.$[selectedItem].updatedAt": now,
			},
		},
		{
			Key: "$inc",
			Value: bson.M{
				"totalPrice": float64(vasItem.Quantity) * vasItem.Price,
			},
		},
	}

	opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{
				{
					Key: "selectedItem._id", Value: itemId,
				},
			},
		},
	})

	res := r.collection.FindOneAndUpdate(context.Background(), filter, update, opts)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (r *cartRepository) UpdateVasItemQuantity(item *domain.Item, vasItem *domain.VasItem, req *models.AddVasItemToItemServiceRequest) error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully updated vasItem to the item", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"items.$[selectedItem].vasItems.$[selectedVasItem].quantity":  vasItem.Quantity,
				"items.$[selectedItem].vasItems.$[selectedVasItem].updatedAt": now,
				"items.$[selectedItem].updatedAt":                             now,
				"updatedAt":                                                   now,
			},
		},
		{
			Key: "$inc",
			Value: bson.M{
				"totalPrice": float64(req.Quantity) * vasItem.Price,
			},
		},
	}

	opts := options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.D{
				{
					Key:   "selectedItem._id",
					Value: item.ID,
				},
			},
			bson.D{
				{
					Key:   "selectedVasItem._id",
					Value: vasItem.ID,
				},
			},
		},
	})

	res := r.collection.FindOneAndUpdate(context.Background(), filter, update, opts)
	if res.Err() != nil {
		panic(res.Err())
	}

	return nil
}

func (r *cartRepository) ResetCart() error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully reset the cart", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"items":      make([]domain.Item, 0),
				"updatedAt":  now,
				"totalPrice": 0,
			},
		},
	}

	res := r.collection.FindOneAndUpdate(context.Background(), filter, update)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (r *cartRepository) RemoveItem(item *domain.Item) error {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully removed item to the cart", "duration", time.Since(now))
	}()

	filter := bson.D{
		{
			Key:   "_id",
			Value: cartID,
		},
	}

	update := bson.D{
		{
			Key: "$pull",
			Value: bson.M{
				"items": bson.M{
					"_id": item.ID,
				},
			},
		},
		{
			Key: "$set",
			Value: bson.M{
				"updatedAt": now,
			},
		},
		{
			Key: "$inc",
			Value: bson.M{
				"totalPrice": float64(item.Quantity) * item.Price * -1,
			},
		},
	}

	res := r.collection.FindOneAndUpdate(context.Background(), filter, update)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (r *cartRepository) FindItemByItemIdFromCart(itemId int) (*domain.Item, error) {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully found item in the cart", "duration", time.Since(now))
	}()

	cart, err := r.GetCart()
	if err != nil {
		return nil, err
	}

	for _, item := range cart.Items {
		if item.ItemId == itemId {
			return item, nil
		}
	}

	return nil, fmt.Errorf("item didn't find in cart")
}

func (r *cartRepository) FindVasItemByVasItemIdFromItem(vasItemId int) (*domain.VasItem, error) {
	now := time.Now()

	defer func() {
		logger.GetLogger().Sugar().Infow("successfully found vasItem in the item", "duration", time.Since(now))
	}()

	cart, err := r.GetCart()
	if err != nil {
		return nil, err
	}

	for _, item := range cart.Items {
		for _, vasItem := range item.VasItems {
			if vasItem.VasItemId == vasItemId {
				return vasItem, nil
			}
		}
	}

	return nil, fmt.Errorf("vasItem didn't find in item")
}
