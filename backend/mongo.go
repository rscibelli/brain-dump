package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPostsMongo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start of GET")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := makeConnection(ctx)
	if err != nil {
		json.NewEncoder(w).Encode("error creating collection: " + err.Error())
		return
	}

	defer close(client, ctx)

	database := client.Database("brain-dump-db")
	collection := database.Collection("posts")

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

func CreatePostMongo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var post Post
	json.Unmarshal(reqBody, &post)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := makeConnection(ctx)
	if err != nil {
		json.NewEncoder(w).Encode("error creating collection: " + err.Error())
		return
	}

	defer close(client, ctx)

	database := client.Database("brain-dump-db")
	collection := database.Collection("posts")

	_, err = collection.InsertOne(context.TODO(), post)
	if err != nil {
		json.NewEncoder(w).Encode("error inserting: " + err.Error())
	}
}

func makeConnection(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://http://192.168.xy.ab:27017/")

	fmt.Println("after client options")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	fmt.Println("after client connect")

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func close(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(context.TODO()); err != nil {
		fmt.Println("failed to disconnect")
	}
	fmt.Println("disconnected from MongoDB!")
}
