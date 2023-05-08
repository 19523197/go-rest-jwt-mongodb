package repository

import (
	"go-jwt-rest-mongodb/repository/userRepository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UserRepo *userRepository.UserRepository
}

func InitRepository(mongo *mongo.Client) *Repository {
	return &Repository{
		UserRepo: userRepository.InitUserRepository(mongo),
	}
}
