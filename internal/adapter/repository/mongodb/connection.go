package mongodb

import (
	"checkout-case/pkg/config"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection() (*mongo.Collection, error) {
	l := logger.GetLogger().Sugar()

	cfg := config.Cfg.MongoDB

	// example: uri := "mongodb://root:example@localhost:27017/?timeoutMS=5000"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?timeoutMS=%d", cfg.User, cfg.Password, cfg.Addr, cfg.Port, cfg.Timeout)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	l.Info("mongodb successfully connected")

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	l.Info("mongodb successfully pinged")

	return client.Database(cfg.Name).Collection(cfg.Collection), nil
}
