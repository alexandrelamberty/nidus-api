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
	ReadUser(ID string) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(ID string) error
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

func (r *repository) ReadUser(ID string) (*domain.User, error) {
	var user *domain.User
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

func (r *repository) DeleteUser(ID string) error {

	fmt.Println(ID)
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

func (r *repository) ListUsers() (*[]domain.User, error) {
	var users []domain.User
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user domain.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}
