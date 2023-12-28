//go:build integration
// +build integration

package mongodb

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
)

var db *mongo.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=root",
		"MONGO_INITDB_ROOT_PASSWORD=password",
	}

	resource, err := pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		uri := fmt.Sprintf("mongodb://root:password@localhost:%s", resource.GetPort("27017/tcp"))

		db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return db.Ping(context.TODO(), nil)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	exitCode := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(exitCode)
}
