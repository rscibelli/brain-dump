package main

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	ws        = "brain-dump-files/posts"
	delimiter = "_____"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var post Post
	json.Unmarshal(reqBody, &post)

	songRequestString := post.Name + delimiter + post.Date + delimiter + post.Title + delimiter + post.Post + "\n"

	f, err := os.OpenFile(ws, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		retMsg := "Unable to add post, error opening file"
		json.NewEncoder(w).Encode(retMsg)
		return
	}

	defer f.Close()

	_, err = f.WriteString(songRequestString)
	if err != nil {
		retMsg := "Unable to add post, error writing to file"
		json.NewEncoder(w).Encode(retMsg)
		return
	}

	retMsg := "Post has been added!"
	json.NewEncoder(w).Encode(retMsg)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	file, err := os.Open(ws)
	if err != nil {
		json.NewEncoder(w).Encode("error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		postArr := strings.Split(scanner.Text(), delimiter)
		post := Post{Name: postArr[0], Date: postArr[1], Title: postArr[2], Post: postArr[3]}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}
