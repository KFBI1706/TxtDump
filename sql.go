package main

import (
	"database/sql"
	"encoding/hex"
	"errors"
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
func createPostDB(post postData) {
	db, err := establishConn()

	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("INSERT INTO text (id, title, text, created_at, editid, views, postperms, hash, salt, key) VALUES ($1, $2, $3, $4, $5, 0, $6, $7, $8, $9); ", post.ID, post.Title, post.Content, time.Now(), post.EditID, post.PostPerms, post.Hash, post.Salt, post.Key)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
func readpostDB(ID int) (postData, error) {
	result := postData{}
	db, err := establishConn()
	err = db.QueryRow("SELECT id, text, title, created_at, views, postperms, salt, key FROM text WHERE id = $1", ID).Scan(&result.ID, &result.Content, &result.Title, &result.Time, &result.Views, &result.PostPerms, &result.Salt, &result.Key)
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
func saveChanges(post postData) error {
	db, err := establishConn()
	if err != nil {
		return err
	}

	if valid := checkPass(post.Hash, post.ID, post.PostPerms); valid {
		_, err = db.Exec("UPDATE text SET title = $1, text = $2 WHERE id = $3;", post.Title, post.Content, post.ID)
		if err != nil {
			return err
		}
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
func deletepost(post postData) error {
	if valid := checkPass(post.Hash, post.ID, post.PostPerms); valid {

		db, _ := establishConn()
		_, err := db.Exec("DELETE FROM text WHERE id = $1", post.ID)
		if err != nil {
			return err
		}
		db.Close()

	} else {
		return errors.New("Wrong Password!")
	}
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

func hexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func getProp(prop string, id int) ([]byte, error) { //todo:encoding parameter
	if prop == "salt" || prop == "hash" {
		var hash string
		db, err := establishConn()
		if err != nil {
			log.Println(err)
		}
		err = db.QueryRow("SELECT "+prop+" FROM text WHERE id = $1", id).Scan(&hash)
		if err != nil {
			log.Println(err)
		}
		db.Close()
		return hexToBytes(hash), err
	}
	return nil, nil
}
