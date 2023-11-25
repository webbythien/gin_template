package service

import (
	"admin-panel/v1/app/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetPostFromMongoDB() ([]model.Post, error) {
	// Get post from mongoDB
	postCollection := model.PostCollection()

	cursor, err := postCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	var posts []model.Post

	for cursor.Next(context.TODO()) {
		var post model.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
