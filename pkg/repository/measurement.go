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

type MeasurementRepository interface {
	ListMeasurements() (*[]domain.Measurement, error)
	CreateMeasurement(user *domain.Measurement) (*domain.Measurement, error)
	ReadMeasurement(ID string) (*domain.Measurement, error)
	UpdateMeasurement(user *domain.Measurement) (*domain.Measurement, error)
	DeleteMeasurement(ID string) error
}

func NewMeasurementRepo(collection *mongo.Collection) MeasurementRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateMeasurement(user *domain.Measurement) (*domain.Measurement, error) {
	user.ID = primitive.NewObjectID()
	//user.CreatedAt = time.Now()
	//user.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadMeasurement(ID string) (*domain.Measurement, error) {
	var user *domain.Measurement
	return user, nil
}

func (r *repository) UpdateMeasurement(user *domain.Measurement) (*domain.Measurement, error) {
	// user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) DeleteMeasurement(ID string) error {
	measurementId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": measurementId})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ListMeasurements() (*[]domain.Measurement, error) {
	var users []domain.Measurement
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user domain.Measurement
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}
