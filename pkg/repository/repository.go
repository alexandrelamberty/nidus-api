package repository

import "go.mongodb.org/mongo-driver/mongo"

type repository struct {
	Collection *mongo.Collection
}
