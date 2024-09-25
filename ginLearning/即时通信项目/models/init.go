// Package models
// author: zfy  date: 2024/9/25
package models

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = InitMongo()
var RDB = InitRedis()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://192.168.1.8:27017"))
	if err != nil {
		log.Println("Connection MongoDB Error:", err)
		return nil
	}
	return client.Database("im")
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
