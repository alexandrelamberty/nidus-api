package repository

import (
	"context"
	"fmt"
	"nidus-server/pkg/domain"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ZoneRepository interface {
	ListZones() (*[]domain.Zone, error)
	CreateZone(user *domain.Zone) (*domain.Zone, error)
	ReadZone(ID string) (*domain.Zone, error)
	UpdateZone(user *domain.Zone) (*domain.Zone, error)
	DeleteZone(ID string) error
}

func NewZoneRepo(collection *mongo.Collection) ZoneRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ListZones() (*[]domain.Zone, error) {
	var users []domain.Zone
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user domain.Zone
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}

func (r *repository) CreateZone(user *domain.Zone) (*domain.Zone, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadZone(ID string) (*domain.Zone, error) {
	var zone *domain.Zone
	return zone, nil
}

func (r *repository) UpdateZone(user *domain.Zone) (*domain.Zone, error) {
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) DeleteZone(ID string) error {
	userID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		return err
	}
	return nil
}
