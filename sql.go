package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//Since the Postgresql Go libary just uses a string for info i just read a file with the private database info in it as a string with this see readme.md for more
func readDBstring(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return "Something went wrong reading the db string. See readme.md for info about this"
	}
	return string(file)
}

func testDBConnection() error {
	dbstring := readDBstring("dbstring")
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("DB connection sucsessfully established")
	db.Close()
	return nil
}

func establishConn() *sql.DB {
	dbstring := readDBstring("dbstring")
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createPostDB(post postresp) {
	db := establishConn()
	postdata, err := db.Exec("INSERT INTO text (pubid, title, text, created_at) VALUES ($1, $2, $3, $4); ", post.PubID, post.Title, post.Content, time.Now())
	if err != nil {
		fmt.Println(err, postdata)
	}
	db.Close()
	log.Println(post)
}

func readpostDB(pubid int) postresp {
	result := postresp{PubID: pubid}
	db := establishConn()
	err := db.QueryRow("SELECT id, text, title, created_at FROM text WHERE pubid = $1", pubid).Scan(&result.ID, &result.Content, &result.Title, &result.Time)
	if err != nil && err == sql.ErrNoRows {
		log.Println(err)
		result.Sucsess = false
		return result
	}
	if err != nil && result.Title == "" {
		result.Title = "No title"
	}
	db.Close()
	result.Sucsess = true
	return result
}
func countPosts() int {
	var count int
	db := establishConn()
	rows, err := db.Query("SELECT COUNT(*) as count FROM text")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	db.Close()
	return count
}

//canned function for now...
func findallposts() postcounter {
	posts := postcounter{}
	db := establishConn()
	rows, err := db.Query("SELECT pubid, title FROM text")
	if err != nil && err == sql.ErrNoRows {
		log.Printf("No posts found most likely this is because the db is not properly setup/has no posts")
		return posts
	}
	var derp int
	var lul string
	for rows.Next() {
		err := rows.Scan(&derp, &lul)
		if err != nil {
			log.Fatal(err)
		}
		//posts = append(posts, postcounter{PostIDs: derp, Titles: lul})
	}
	log.Println(posts)
	db.Close()
	return posts
}
