package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type postcounter struct {
	Count   int      `json:"Count"`
	PostIDs []int    `json:"ID"`
	Titles  []string `json:"Titles"`
}

//Since the Postgresql Go libary just uses a string for info i just read a file with the private database info in it as a string with this see readme.md for more
func readDBstring(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(file), nil
}
func testDBConnection() error {
	dbstring, err := readDBstring("dbstring")
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("DB connection sucsessfully established")
	db.Close()
	return nil
}

//establishConn() creates the DB Connnection Remember to always close these u dumbus
func establishConn() (*sql.DB, error) {
	dbstring, err := readDBstring("dbstring")
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func createPostDB(post postdata) {
	db, err := establishConn()
	_, err = db.Exec("INSERT INTO text (id, title, text, created_at, editid, views) VALUES ($1, $2, $3, $4, $5, 0); ", post.ID, post.Title, post.Content, time.Now(), post.EditID)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
func readpostDB(ID int) (postdata, error) {
	result := postdata{ID: ID}
	db, err := establishConn()
	err = db.QueryRow("SELECT id, text, title, created_at, views FROM text WHERE id = $1", ID).Scan(&result.ID, &result.Content, &result.Title, &result.Time, &result.Views)
	db.Close()

	if err != nil && err == sql.ErrNoRows {
		log.Println(err)
		return result, err
	}
	if err != nil && result.Title == "" {
		result.Title = ""
	}
	return result, err
}
func checkedid(post postdata) error {
	db, err := establishConn()
	var edid string
	err = db.QueryRow("SELECT editid FROM text WHERE id = $1", post.ID).Scan(&edid)
	if err != nil {
		log.Println(err)
	}
	if edid != post.EditID {
		err = errors.New("Provided ID is not the same as in DB")
		return err
	}
	db.Close()
	return nil
}
func saveChanges(post postdata) error {
	db, err := establishConn()
	checkPass(post.EditID, post.ID)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE text SET title = $1, text = $2 WHERE id = $3;", post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
func countPosts() int {
	var count int
	db, err := establishConn()
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
func deletepost(post postdata) error {
	db, err := establishConn()
	_, err = db.Exec("DELETE FROM text WHERE id = $1 AND editid = $2", post.ID, post.EditID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
func incrementViewCounter(id int) error {
	db, err := establishConn()
	_, err = db.Exec("UPDATE text SET views = views + 1 WHERE id = $1", id)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
func checkForDuplicateID(id int) bool {
	db, err := establishConn()
	if err != nil {
		log.Println(err)
	}
	res := db.QueryRow("SELECT id FROM text WHERE id = $1", id).Scan(id)
	db.Close()
	if res != sql.ErrNoRows {
		return false
	}
	return true
}
func getHashedPS(id int) []byte {
	var ps []byte
	db, err := establishConn()
	if err != nil {
		log.Println(err)
	}
	err = db.QueryRow("SELECT editid FROM text WHERE id = $1").Scan(ps)
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return ps
}
