package model

type User struct {
	Id   int    `bson:"user_id"`
	Name string `bson:"name"`
}
