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

type UserRepository interface {
	ListUsers() (*[]domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id string) error
}


func NewUserRepo(collection *mongo.Collection) UserRepository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	//user.CreatedAt = time.Now()
	//user.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadUser(id string) (*domain.User, error) {
	var user *domain.User
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": userId}
	err = r.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) UpdateUser(user *domain.User) (*domain.User, error) {
	// user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) DeleteUser(id string) error {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ListUsers() (*[]domain.User, error) {
	var users []domain.User
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("ListUsers", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user domain.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}
