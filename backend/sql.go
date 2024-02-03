package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (*sql.DB, error) {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "braindump", "root", "localhost", "3306", "brain-dump-db"))
	if err != nil {
		return nil, err
	}

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreatePostSQL(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var post Post
	json.Unmarshal(reqBody, &post)

	db, err := getConnection()
	if err != nil {
		json.NewEncoder(w).Encode("error creating connection: " + err.Error())
		return
	}
	defer db.Close()
	// Prepare the SQL statement for inserting a row
	stmt, err := db.Prepare("INSERT INTO your_table_name (name, date, post) VALUES (?, ?, ?)")
	if err != nil {
		json.NewEncoder(w).Encode("error preparing sql: " + err.Error())
		return
	}
	defer stmt.Close()

	// Execute the SQL statement with parameters
	_, err = stmt.Exec(post.Name, post.Date, post.Post)
	if err != nil {
		json.NewEncoder(w).Encode("error executing sql: " + err.Error())
		return
	}

	fmt.Println("Row inserted successfully.")
}

func GetPostsSQL(w http.ResponseWriter, r *http.Request) {
	db, err := getConnection()
	if err != nil {
		json.NewEncoder(w).Encode("error creating connection: " + err.Error())
		return
	}
	defer db.Close()
	// Query all rows from the table
	rows, err := db.Query("SELECT * FROM your_table_name")
	if err != nil {
		json.NewEncoder(w).Encode("error querying: " + err.Error())
		return
	}
	defer rows.Close()

	var posts []Post

	// Iterate through the result set
	for rows.Next() {
		var postRaw []byte
		if err := rows.Scan(&postRaw); err != nil {
			json.NewEncoder(w).Encode("error scanning: " + err.Error())
			return
		}
		var post Post
		json.Unmarshal(postRaw, &post)
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}
