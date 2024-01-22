package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/post", CreatePost).Methods("POST")
	router.HandleFunc("/api/posts", GetPosts).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Starting server at port 2081")
	if err := http.ListenAndServe(":2081", handler); err != nil {
		log.Fatal(err)
	}
}

// env GOOS=linux GOARCH=arm GOARM=7 go build
