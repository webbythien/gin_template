package database

import (
	"admin-panel/v1/pkg/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func NewContext() (context.Context, context.CancelFunc) {
	// return context.WithTimeout(context.Background(), 10*time.Second)
	return context.TODO(), nil
}

func init() {
	initDatabase()
}

func initDatabase() {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/?retryWrites=true&w=majority", config.DbType, config.DbUser, config.DbPass, config.DbHost, config.DbPort)
	fmt.Printf("URL MongoDB %s\n", url)
	ctx, _ := NewContext()
	var err error
	clientOptions := options.Client().ApplyURI(url)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("[MONGO_DB] Cannot connect to mongoDB", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Cannot ping MongoDB: %v", err)
	}
	log.Println("[MONGO_DB] connect successfully")
}

func Shutdown() {
	ctx, _ := NewContext()
	if client != nil {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Cannot disconnect MongoDB: %v", err)
		}
	}
	log.Println("[MONGO_DB] disconnect successfully")
}

func GetCollection(col string) *mongo.Collection {
	if client == nil {
		initDatabase()
	}
	return client.Database(config.DbName).Collection(col)
}
