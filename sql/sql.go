package sql

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/jinzhu/gorm"
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
		return "", err
	}
	return string(file), nil
}

//TestDBConnection Just pings the DB
func TestDBConnection() error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	var postData model.PostData
	if db.HasTable(&postData) != true {
		return errors.New("Table does not exist. Run with -setupdb to fix this")
	}
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
	return db, nil
}

//CreatePostDB registers the post in the DB
func CreatePostDB(post model.PostData) error {
	db, err := EstablishConn()
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//ReadPostDB gets postdata from DB
func ReadPostDB(ID int) (model.PostData, error) {
	var result model.PostData
	db, err := EstablishConn()
	defer db.Close()
	err = db.First(&result, ID).Error
	return result, err
}

//SaveChanges registers edits in DB
func SaveChanges(post model.PostData) error {
	db, err := EstablishConn()
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.Model(&model.PostData{}).UpdateColumns(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//CountPosts runs SQL count on DB
func CountPosts() int {
	db, err := EstablishConn()
	if err != nil {
		log.Println("sql ", err)
	}
	defer db.Close()
	var count int
	err = db.Model(&model.PostData{}).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return count
}

//PostMetas returns some data used for index overview
func PostMetas() (model.PostCounter, error) {
	postMeta := []model.PostData{}
	posts := model.PostCounter{Count: CountPosts()}
	db, err := EstablishConn()
	defer db.Close()
	if err != nil {
		return posts, err
	}
	err = db.Find(&postMeta, &model.PostData{}).Error
	if err != nil {
		log.Println(err)
	}
	posts.Meta = postMeta
	return posts, err
}

//DeletePost deletes post in DB
func DeletePost(post model.PostData) error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	err = db.Delete(&post).Error
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
	err = db.Exec("UPDATE post_data SET views = views + 1 WHERE id = $1", id).Error
	if err != nil {
		log.Println(err)
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
	if db.NewRecord(model.PostData{ID: id}) {
		return false
	}
	db.Close()

	return true
}

//GetProp gets the requested hash from the DB
func GetProp(prop string, id int) ([]byte, error) { //todo:encoding parameter
	if prop == "salt" || prop == "hash" {
		var hash []string
		db, err := EstablishConn()
		if err != nil {
			log.Println(err)
		}
		err = db.First(&model.PostData{}, id).Pluck(prop, &hash).Error
		if err != nil {
			return nil, err
		}
		db.Close()
		return HexToBytes(hash[0]), err
	}
	return nil, nil
}

/*SetupDB is used to setup the database
no input arguments
returns error*/
func SetupDB() error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	var postData model.PostData
	if !db.HasTable(&postData) {
		db.AutoMigrate(&postData)
	}
	return nil
}

/*ClearOutDB is used to clear a table
no input arguments
returns error*/
func ClearOutDB() error {
	db, err := EstablishConn()
	if err != nil {
		return err
	}
	var postData model.PostData
	db.DropTableIfExists(&postData)
	return nil
}
