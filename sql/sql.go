package sql

import (
	"database/sql"
	"encoding/hex"
	"io/ioutil"
	"log"
	"time"

	"github.com/KFBI1706/Txtdump/model"
	//Import Postgres Libary
	_ "github.com/lib/pq"
)

// HexToBytes decodes the inputted hex string
func HexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

//ReadDBstring Returns the file in filename as a string.
func ReadDBstring(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(file), nil
}

//TestDBConnection Just pings the DB
func TestDBConnection() error {
	dbstring, err := ReadDBstring("dbstring")
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

//EstablishConn creates the DB Connnection
func EstablishConn() (*sql.DB, error) {
	dbstring, err := ReadDBstring("dbstring")
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

//CreatePostDB registers the post in the DB
func CreatePostDB(post model.PostData) error {
	db, err := EstablishConn()
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO text (id, title, text, created_at, editid, views, postperms, hash, salt, key) VALUES ($1, $2, $3, $4, $5, 0, $6, $7, $8, $9); ", post.ID, post.Title, post.Content, time.Now(), post.EditID, post.PostPerms, post.Hash, post.Salt, post.Key)
	if err != nil {
		return err
	}
	return nil
}

//ReadPostDB gets postdata from DB
func ReadPostDB(ID int) (model.PostData, error) {
	result := model.PostData{}
	db, err := EstablishConn()
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

//SaveChanges registers edits in DB
func SaveChanges(post model.PostData) error {
	db, err := EstablishConn()
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

//CountPosts runs SQL count on DB
func CountPosts() int {
	var count int
	db, err := EstablishConn()
	err = db.QueryRow("SELECT COUNT(*) as count FROM text").Scan(&count)
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return count
}

//PostMetas returns some data used for index overview
func PostMetas() (model.PostCounter, error) {
	posts := model.PostCounter{}
	db, err := EstablishConn()
	if err != nil {
		return posts, err
	}
	err = db.QueryRow("SELECT COUNT(*) AS count FROM text").Scan(&posts.Count)
	if err != nil {
		return posts, err
	}
	rows, err := db.Query("SELECT id, title, views, postperms FROM text LIMIT 20")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var meta model.PostMeta
		rows.Scan(&meta.PostID, &meta.Title, &meta.Views, &meta.PostPerms)
		posts.Meta = append(posts.Meta, meta)
	}
	db.Close()
	return posts, err
}

//DeletePost deletes post in DB
func DeletePost(post model.PostData) error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM text WHERE id = $1", post.ID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

//IncrementViewCounter increments the viewcounter in DB
func IncrementViewCounter(id int) error {
	db, err := EstablishConn()
	_, err = db.Exec("UPDATE text SET views = views + 1 WHERE id = $1", id)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

//CheckForDuplicateID Checks if ID is already used for a post in DB
func CheckForDuplicateID(id int) bool {
	db, err := EstablishConn()
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

//GetProp gets the requested hash from the DB
func GetProp(prop string, id int) ([]byte, error) { //todo:encoding parameter
	if prop == "salt" || prop == "hash" {
		var hash string
		db, err := EstablishConn()
		if err != nil {
			log.Println(err)
		}
		err = db.QueryRow("SELECT "+prop+" FROM text WHERE id = $1", id).Scan(&hash)
		if err != nil {
			log.Println(err)
		}
		db.Close()
		return HexToBytes(hash), err
	}
	return nil, nil
}
