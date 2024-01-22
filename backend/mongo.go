package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	collection, err := makeCollection()
	if err != nil {
		json.NewEncoder(w).Encode("error creating collection: " + err.Error())
		return
	}

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		json.NewEncoder(w).Encode("error finding in collection: " + err.Error())
		return
	}

	var posts []Post
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var entry Post
		if err := cursor.Decode(&entry); err != nil {
			json.NewEncoder(w).Encode("error decoding entry: " + err.Error())
			return
		}
		posts = append(posts, entry)
	}

	if err := cursor.Err(); err != nil {
		json.NewEncoder(w).Encode("error with cursor: " + err.Error())
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var post Post
	json.Unmarshal(reqBody, &post)

	collection, err := makeCollection()
	if err != nil {
		json.NewEncoder(w).Encode("error creating collection: " + err.Error())
		return
	}

	_, err = collection.InsertOne(context.TODO(), post)
	if err != nil {
		json.NewEncoder(w).Encode("error inserting: " + err.Error())
	}
}

func makeCollection() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	// Close the connection when done
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Println("failed to disconnect")
		}
	}()

	database := client.Database("brain-dump-db")

	// Access a collection within the database
	return database.Collection("posts"), nil
}
