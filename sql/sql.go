package sql

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/KFBI1706/Txtdump/model"
	//Import Postgres Libary
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	return nil
}

//EstablishConn creates the DB Connnection
func EstablishConn() (*gorm.DB, error) {
	dbstring, err := ReadDBstring("dbstring")
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open("postgres", dbstring)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.PostData{})
	return db, nil
}
func EstablishConnOld() (*sql.DB, error) {
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
	return db, err
}

//CreatePostDB registers the post in the DB
func CreatePostDB(post model.PostData) error {
	db, err := EstablishConn()
	defer db.Close()
	if err != nil {
		return err
	}
	db.Create(&post)
	if err != nil {
		return err
	}
	return nil
}

//ReadPostDB gets postdata from DB
func ReadPostDB(ID int) (model.PostData, error) {
	result := model.PostData{}
	db, err := EstablishConn()
	db.First(&result, ID)
	db.Close()
	return result, err
}

//SaveChanges registers edits in DB
func SaveChanges(post model.PostData) error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	db.Where("ID = $1", post.ID).Updates(post)
	db.Close()
	return nil
}

//CountPosts runs SQL count on DB
func CountPosts() int {
	var count int
	db, err := EstablishConn()
	db.First("SELECT COUNT(*) as count FROM post_data").Scan(&count)
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
	db.First("SELECT COUNT(*) AS count FROM post_data").Scan(&posts.Count)
	if err != nil {
		return posts, err
	}
	/*	db.Query("SELECT id, title, views, postperms FROM post_data LIMIT 20")
		if err != nil {
			return posts, err
		}
		for rows.Next() {
			var meta model.PostMeta
			rows.Scan(&meta.PostID, &meta.Title, &meta.Views, &meta.PostPerms)
			posts.Meta = append(posts.Meta, meta)
		}
		db.Close()
	*/return posts, err
}

//DeletePost deletes post in DB
func DeletePost(post model.PostData) error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	db.Delete(&post.ID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

//IncrementViewCounter increments the viewcounter in DB
func IncrementViewCounter(id int) error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	db.Exec("UPDATE text SET views = views + 1 WHERE id = $1", id)
	defer db.Close()
	return nil
}

//CheckForDuplicateID Checks if ID is already used for a post in DB
func CheckForDuplicateID(id int) bool {
	db, err := EstablishConn()
	if err != nil {
		log.Println(err)
	}
	db.First(&id)
	db.Close()

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
		db.First("SELECT "+prop+" FROM post_data WHERE id = $1", id).Scan(&hash)
		if err != nil {
			log.Println(err)
		}
		db.Close()
		return HexToBytes(hash), err
	}
	return nil, nil
}

/*SetupDB is used to setup the database
no input arguments
returns error*/
func SetupDB() error {
	db, err := EstablishConnOld()
	if err != nil {
		return err
	}
	sql, err := ReadDBstring("sql/db.sql")
	if err != nil {
		return err
	}
	res, err := db.Exec(sql)
	fmt.Println(res)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

/*ClearOutDB is used to clear a table
no input arguments
returns error*/
func ClearOutDB() error {
	db, err := EstablishConnOld()
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE text, text_test;")
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}
