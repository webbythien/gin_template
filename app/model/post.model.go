package model

import (
	"admin-panel/v1/platform/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	postCollection = "posts"
)

type Post struct {
	ID     int32  `json:"id"`
	UserId int32  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func PostCollection() *mongo.Collection {
	return database.GetCollection(postCollection)
}
