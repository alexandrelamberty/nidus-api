package infrastructure

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoDBConnection() (*mongo.Database, error) {
	var uri = os.Getenv("DATABASE_URI")

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri).SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
		// panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database("nidus")
	return db, nil
}
