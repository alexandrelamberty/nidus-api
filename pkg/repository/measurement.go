package repository

import (
	"context"
	"fmt"
	"nidus-server/internal/requests"
	"nidus-server/pkg/domain"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MeasurementRepository interface {
	ListMeasurements() (*[]domain.Measurement, error)
	CreateMeasurement(user *requests.CreateMeasurementRequest) (*domain.Measurement, error)
	ReadMeasurement(ID string, sensorType string, timestamp string) (*domain.Measurement, error)
}

func NewMeasurementRepo(
	temperature *mongo.Collection,
	humidity *mongo.Collection,
	pressure *mongo.Collection) MeasurementRepository {
	return &repository{
		Collection: humidity,
	}
}

func (r *repository) ListMeasurements() (*[]domain.Measurement, error) {
	var measurements []domain.Measurement

	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var measurement domain.Measurement
		_ = cursor.Decode(&measurement)
		measurements = append(measurements, measurement)
	}
	return &measurements, nil
}

func (r *repository) LastMeasurement(device_id string) (*[]domain.Measurement, error) {
	var measurements []domain.Measurement
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetLimit(1)
	objID, _ := primitive.ObjectIDFromHex("6323ad8d2812cd11dc72a05a")
	cursor, err := r.Collection.Find(context.TODO(), bson.D{{"metadata.device_id", objID}}, findOptions)
	fmt.Println(cursor)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var measurement domain.Measurement
		_ = cursor.Decode(&measurement)
		measurements = append(measurements, measurement)
	}
	return &measurements, nil
}

func (r *repository) CreateMeasurement(measurement *requests.CreateMeasurementRequest) (*domain.Measurement, error) {
	_, err := r.Collection.InsertOne(context.Background(), measurement)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *repository) ReadMeasurement(ID string, sensorType string, timestamp string) (*domain.Measurement, error) {
	var measurement *domain.Measurement
	fmt.Println(ID, sensorType, timestamp)
	return measurement, nil
}
