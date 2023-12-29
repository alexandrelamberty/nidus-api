package repository

import (
	"context"
	"fmt"
	"log"
	"nidus-server/pkg/domain"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CapabilityRepository interface {
	ListCapabilities() (*[]domain.Capability, error)
	CreateCapability(capability *domain.Capability) (*domain.Capability, error)
	ReadCapability(ID string) (*domain.Capability, error)
	UpdateCapability(capability *domain.Capability) (*domain.Capability, error)
	DeleteCapability(ID string) error
}

func NewCapabilityRepo(collection *mongo.Collection) CapabilityRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) ListCapabilities() (*[]domain.Capability, error) {
	fmt.Println("ListCapabilities")
	var capabilities []domain.Capability
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("CapabilityRepository", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var capability domain.Capability
		err = cursor.Decode(&capability)
		if err != nil {
			log.Fatal(err)
		}
		capabilities = append(capabilities, capability)
	}
	return &capabilities, nil
}

func (r *repository) CreateCapability(capability *domain.Capability) (*domain.Capability, error) {
	capability.ID = primitive.NewObjectID()
	//capability.CreatedAt = time.Now()
	//capability.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), capability)
	if err != nil {
		return nil, err
	}
	return capability, nil
}

func (r *repository) ReadCapability(ID string) (*domain.Capability, error) {
	var capability *domain.Capability
	return capability, nil
}

func (r *repository) UpdateCapability(capability *domain.Capability) (*domain.Capability, error) {
	// capability.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": capability.ID}, bson.M{"$set": capability})
	if err != nil {
		return nil, err
	}
	return capability, nil
}

func (r *repository) DeleteCapability(ID string) error {
	capabilityID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": capabilityID})
	if err != nil {
		return err
	}
	return nil
}
