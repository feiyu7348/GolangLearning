// Package test
// author: zfy  date: 2024/9/25
package test

import (
	"context"
	"fmt"
	"im/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestfindOne() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "root",
		Password: "123456",
	}).ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("im")
	ub := new(models.UserBasic)
	err = db.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ub ==> ", ub)
}
