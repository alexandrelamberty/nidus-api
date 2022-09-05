package infrastructure

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FIXME: Use environment variables
func MongoDBConnection() (*mongo.Database, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf(err.Error())
		//cancel()
		return nil, err
	}
	db := client.Database("nidus")
	return db, nil
}
