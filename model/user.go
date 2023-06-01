package model

type User struct {
	Id   int    `bson:"id, omitempty"`
	Name string `bson:"name"`
}
