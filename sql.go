package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type postcounter struct {
	Count int        `json:"Count"`
	Meta  []postmeta `json:"Meta"`
}
type postmeta struct {
	PostID int    `json:"ID"`
	Title  string `json:"Title"`
	Views  int    `json:"View"`
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
	post.EditID, err = securePass(post.EditID)
	if err != nil {
		log.Println(err)
	}
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
func saveChanges(post postdata) error {
	db, err := establishConn()
	if err != nil {
		return err
	}
	err = checkPass(post.EditID, post.ID)
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
	err = db.QueryRow("SELECT COUNT(*) as count FROM text").Scan(&count)
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return count
}
func postMeta() (postcounter, error) {
	posts := postcounter{}
	db, err := establishConn()
	if err != nil {
		return posts, err
	}
	err = db.QueryRow("SELECT COUNT(*) AS count FROM text").Scan(&posts.Count)
	if err != nil {
		return posts, err
	}
	rows, err := db.Query("SELECT id, title, views FROM text LIMIT 20")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var meta postmeta
		rows.Scan(&meta.PostID, &meta.Title, &meta.Views)
		posts.Meta = append(posts.Meta, meta)
	}
	db.Close()
	return posts, err
}
func deletepost(post postdata) error {
	err := checkPass(post.EditID, post.ID)
	if err != nil {
		return err
	}
	db, err := establishConn()
	_, err = db.Exec("DELETE FROM text WHERE id = $1", post.ID)
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
	var hash string
	db, err := establishConn()
	if err != nil {
		log.Println(err)
	}
	err = db.QueryRow("SELECT editid FROM text WHERE id = $1", id).Scan(&hash)
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return []byte(hash)
}
