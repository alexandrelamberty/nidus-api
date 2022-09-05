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
	ReadDevice(ID string) (*domain.Device, error)
	UpdateDevice(user *domain.Device) (*domain.Device, error)
	DeleteDevice(ID string) error
}

func NewDeviceRepo(collection *mongo.Collection) DeviceRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ListDevices() (*[]domain.Device, error) {
	var devices []domain.Device
	// opts := options.FindOne().SetSort(bson.D{{"age", 1}})
	//cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	var device domain.Device
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "zones"}, {"localField", "zone"}, {"foreignField", "_id"}, {"as", "zone"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$zone"}, {"preserveNullAndEmptyArrays", false}}}}
	cursor, err := r.Collection.Aggregate(context.Background(), mongo.Pipeline{lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.Background(), &devices); err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		_ = cursor.Decode(&device)
		devices = append(devices, device)
	}
	return &devices, nil
}

func (r *repository) CreateDevice(device *domain.Device) (*domain.Device, error) {
	device.ID = primitive.NewObjectID()
	fmt.Printf("%+v\n", device)
	//user.CreatedAt = time.Now()
	//user.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *repository) ReadDevice(ID string) (*domain.Device, error) {
	var device *domain.Device
	deviceId, err := primitive.ObjectIDFromHex(ID)
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

func (r *repository) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	//user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": device.ID}, bson.M{"$set": device})
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (r *repository) DeleteDevice(ID string) error {
	deviceId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": deviceId})
	if err != nil {
		return err
	}
	return nil
}
