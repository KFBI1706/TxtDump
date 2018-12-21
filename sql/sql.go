package sql

import (
	"encoding/hex"
	"errors"
	"log"

	"github.com/KFBI1706/TxtDump/config"
	"github.com/KFBI1706/TxtDump/model"
	"github.com/jinzhu/gorm"
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
	var post model.Post
	var data model.Data
	if config.DB.HasTable(&post) != true || config.DB.HasTable(&data) != true {
		return errors.New("Table does not exist. Run with -setupdb to fix this")
	}
	return nil
}

//CreatePostDB registers the post in the DB
func CreatePostDB(post model.Post) error {
	err := config.DB.Create(&post).Error
	if err != nil {
		return err
	}
	if err = SaveChanges(post); err != nil {
		return err
	}
	log.Println(ReadPostDB(post.ID))
	return nil
}

//ReadPostDB gets postdata from DB
func ReadPostDB(ID int) (model.Post, error) {
	var result model.Post
	err := config.DB.Preload("Data").Preload("Meta").First(&result, ID).Error
	return result, err
}

//SaveChanges registers edits in DB
func SaveChanges(post model.Post) (err error) {
	if err = config.DB.Debug().Model(&model.Data{}).UpdateColumns(&post.Data).Error; err != nil {
		return err
	}
	if err = config.DB.Debug().Model(&model.Crypto{}).UpdateColumns(&post.Crypto).Error; err != nil {
		return err
	}
	return err
}

//CountPosts runs SQL count on DB
func CountPosts() int {
	var count int
	err := config.DB.Model(&model.Post{}).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return count
}

//PostMetas returns some metadata used for index overview
func PostMetas() (metas []model.Meta, err error) {
	//Use pluck
	var posts []model.Post
	if err = config.DB.Debug().Preload("Meta").Find(&posts).Error; err != nil {
		log.Println(err)
		return metas, err
	}
	metas = make([]model.Meta, len(posts))
	for i := range posts {
		metas[i] = posts[i].Meta
	}
	log.Println(metas)
	return metas, err
}

//PostDatas returns some data used for index overview
func PostDatas() (datas []model.Data, err error) {
	//Use pluck
	var posts []model.Post
	if err = config.DB.Debug().Preload("Data").Find(&posts).Error; err != nil {
		log.Println(err)
		return datas, err
	}
	datas = make([]model.Data, len(posts))
	for i := range posts {
		datas[i] = posts[i].Data
	}
	return datas, err
}

//DeletePost deletes post in DB
func DeletePost(post model.Post) error {
	err := config.DB.Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}

//IncrementViewCounter increments the viewcounter in DB
func IncrementViewCounter(id int) error {
	err := config.DB.Model(&model.Meta{}).Update("views", gorm.Expr("views + 1")).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//CheckForDuplicateID Checks if ID is already used for a post in DB
func CheckForDuplicateID(id int) bool {
	if config.DB.NewRecord(model.PostNew{Post: model.Post{ID: id}}) {
		return false
	}
	return true
}

//GetProp gets the requested hash from the DB
func GetProp(prop string, id int) ([]byte, error) { //todo:encoding parameter
	if prop == "salt" || prop == "hash" {
		var s string
		err := config.DB.Debug().Model(&model.Crypto{}).Select(prop).Where(&model.Crypto{PostID: id}).Row().Scan(&s)
		log.Println(s)
		if err != nil {
			return nil, err
		}
		return HexToBytes(s), err
	}
	return nil, nil
}

/*SetupDB is used to setup the database
no input arguments
returns error*/
func SetupDB() error {
	return config.DB.Debug().AutoMigrate(&model.Post{}, &model.Data{}, model.Meta{}, model.Crypto{}, model.Markdown{}, model.Edit{}).Error
}

/*ClearOutDB is used to clear a table
no input arguments
returns error*/
func ClearOutDB() error {
	var post model.Post
	config.DB.DropTableIfExists(&post)
	return nil
}
