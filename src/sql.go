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
		log.Fatal(err)
	}
	return string(file)
}

func testDBConnection() {
	dbstring := readDBstring("dbstring")
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connetcion sucsessfully established")
	db.Close()
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
	postdata, err := db.Exec("INSERT INTO text (pubid, text, created_at) VALUES ($1, $2, $3); ", post.PubID, post.Content, time.Now())
	if err != nil {
		fmt.Println(err, postdata)
	}
	db.Close()
}
