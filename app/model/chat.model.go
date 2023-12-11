package model

import (
	"admin-panel/v1/platform/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	chatCollection = "chats"
)

type Chat struct {
	ID       int32  `bson:"_id"`
	Message  string `json:"message"`
	RoomID   string `json:"room_id"`
	CreateAt string `json:"createAt"`
	UserId   string `bson:"user_id"`
}

func ChatColleChatction() *mongo.Collection {
	return database.GetCollection(chatCollection)
}

type SendMessageReq struct {
	Message string `json:"message"  binding:"required"`
	RoomID  string `json:"room_id"  binding:"required"`
	UserId  string `json:"user_id" bson:"user_id"  binding:"required"`
}
