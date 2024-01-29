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

type DeviceRepository interface {
	ListDevices() (*[]domain.Device, error)
	CreateDevice(user *domain.Device) (*domain.Device, error)
	ReadDevice(id string) (*domain.Device, error)
	UpdateDevice(id string, user *domain.Device) (*domain.Device, error)
	DeleteDevice(id string) error
}

func NewDeviceRepo(collection *mongo.Collection) DeviceRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ListDevices() (*[]domain.Device, error) {
	var devices []domain.Device

	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var device domain.Device
		if err := cursor.Decode(&device); err != nil {
			return nil, err
		}
		fmt.Println(device)
		devices = append(devices, device)
	}

	return &devices, nil
}

func (r *repository) CreateDevice(device *domain.Device) (*domain.Device, error) {
	device.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *repository) ReadDevice(id string) (*domain.Device, error) {
	var device *domain.Device
	deviceId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": deviceId}
	err = r.Collection.FindOne(context.Background(), filter).Decode(&device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *repository) UpdateDevice(id string, device *domain.Device) (*domain.Device, error) {
	update := bson.M{"$set": device}
	filter := bson.M{"_id": id}
	result, err := r.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		fmt.Println("No documents were matched and updated.")
		err := mongo.ErrNoDocuments
		return nil, err
	}
	return device, nil
}

func (r *repository) DeleteDevice(id string) error {
	deviceId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": deviceId})
	if err != nil {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}
