package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type studentMongoDB struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connectMongoDB() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("db_novgo"), nil
}

func insertMongoDB() {
	db, err := connectMongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, studentMongoDB{"John Doe", 4})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, studentMongoDB{"Ethan Hunt", 5})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert to MongoDB success!")
}

// read from database
func find() {
	db, err := connectMongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Ethan Hunt"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]studentMongoDB, 0)
	for csr.Next(ctx) {
		var row studentMongoDB
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name :", result[0].Name)
		fmt.Println("Grade : ", result[0].Grade)
	}
}

func main() {
	fmt.Println("No SQL using MongoDB")
	fmt.Printf("\n")
	insertMongoDB()
	find()
}
