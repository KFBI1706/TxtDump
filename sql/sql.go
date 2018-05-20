package sql

import (
	"encoding/hex"
	"errors"
	"log"

	"github.com/KFBI1706/TxtDump/config"
	"github.com/KFBI1706/TxtDump/model"
	//Import Postgres Library
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

//TestDBConnection Just pings the DB
func TestDBConnection() error {
	var postData model.PostData
	if config.DB.HasTable(&postData) != true {
		return errors.New("Table does not exist. Run with -setupdb to fix this")
	}
	return nil
}

//CreatePostDB registers the post in the DB
func CreatePostDB(post model.PostData) error {
	err := config.DB.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//ReadPostDB gets postdata from DB
func ReadPostDB(ID int) (model.PostData, error) {
	var result model.PostData
	err := config.DB.First(&result, ID).Error
	return result, err
}

//SaveChanges registers edits in DB
func SaveChanges(post model.PostData) error {
	err := config.DB.Model(&model.PostData{}).UpdateColumns(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//CountPosts runs SQL count on DB
func CountPosts() int {
	var count int
	err := config.DB.Model(&model.PostData{}).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return count
}

//PostMetas returns some data used for index overview
func PostMetas() (model.PostCounter, error) {
	postMeta := []model.PostData{}
	posts := model.PostCounter{Count: CountPosts()}
	err := config.DB.Find(&postMeta, &model.PostData{}).Error
	if err != nil {
		log.Println(err)
	}
	posts.Meta = postMeta
	return posts, err
}

//DeletePost deletes post in DB
func DeletePost(post model.PostData) error {
	err := config.DB.Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//IncrementViewCounter increments the viewcounter in DB
func IncrementViewCounter(id int) error {
	err := config.DB.Exec("UPDATE post_data SET views = views + 1 WHERE id = $1", id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//CheckForDuplicateID Checks if ID is already used for a post in DB
func CheckForDuplicateID(id int) bool {
	if config.DB.NewRecord(model.PostData{ID: id}) {
		return false
	}
	return true
}

//GetProp gets the requested hash from the DB
func GetProp(prop string, id int) ([]byte, error) { //todo:encoding parameter
	if prop == "salt" || prop == "hash" {
		var hash []string
		err := config.DB.First(&model.PostData{}, id).Pluck(prop, &hash).Error
		if err != nil {
			return nil, err
		}
		return HexToBytes(hash[0]), err
	}
	return nil, nil
}

/*SetupDB is used to setup the database
no input arguments
returns error*/
func SetupDB() error {
	var postData model.PostData
	if !config.DB.HasTable(&postData) {
		config.DB.AutoMigrate(&postData)
	}
	return nil
}

/*ClearOutDB is used to clear a table
no input arguments
returns error*/
func ClearOutDB() error {
	var postData model.PostData
	config.DB.DropTableIfExists(&postData)
	return nil
}
