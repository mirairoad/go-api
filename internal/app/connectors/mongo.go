package connectors

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mirairoad/goApi/internal/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database

func InitDb(cfg *config.Environment) *mongo.Database {
	uri := cfg.MongoUri
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(cfg.DbName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully!")

	return DB
}
