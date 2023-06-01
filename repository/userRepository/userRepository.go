package userRepository

import (
	"context"
	"go-jwt-rest-mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Database
}

func InitUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		DB: db.Database("primary"),
	}
}

func (r *UserRepository) GetUser() ([]model.User, error) {
	var users []model.User

	collection := r.DB.Collection("user")
	results, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	if err := results.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) InsertUser(user model.User) error {
	collection := r.DB.Collection("user")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}
