package model

import (
	"html/template"

	"github.com/jinzhu/gorm"
)

//M Generic map[string]interface{} to decode into
type M map[string]interface{}

//Post is the "master" struct
type Post struct {
	ID       int `json:"ID" gorm:"primary_key"`
	Data     Data
	Meta     Meta
	Edit     Edit     `json:"-"`
	Crypto   Crypto   `json:"-"`
	Markdown Markdown `json:"-"`
}

//Data contains all the info related to the posts
type Data struct {
	gorm.Model
	PostID    int
	Title     string `json:"Title" gorm:"Column:title"`
	PostPerms int    `json:"PostPerms,string" gorm:"Column:postperms"`
	Content   string `json:"Content" gorm:"Column:text"`
}

//Crypto contains the cryptographic aspect of the post
type Crypto struct {
	gorm.Model
	PostID   int
	Hash     string `json:"Password" gorm:"Column:hash"`
	Salt     string `json:"Salt" gorm:"Column:salt"`
	AuthHash string `json:"authHash" gorm:"-"`
	Key      string `json:"Key" gorm:"Column:key"`
}

//Meta Contains most of the metadata views timestamps etc
type Meta struct {
	gorm.Model
	PostID int
	Views  int `json:"Views" gorm:"Column:views"`
}

//Markdown Contains the data used for rendering the post on the HTML frontend
type Markdown struct {
	MD      template.HTML `json:"" gorm:"-"`
	TitleMD template.HTML `json:"" gorm:"-"`
	IMG     template.HTML `json:"" gorm:"-"`
}

//PostNew is used for registering a new post
type PostNew struct {
	Post
	//ID int `json:"ID" gorm:"Column:id;primary_key"`
	//edit PostEdit
	//meta PostMeta
	//data PostData
}

//Edit is used for editing a post
type Edit struct {
	EditID int `json:"EditID" gorm:"Column:editid"`
}

//PostCounter is used on index to provide some metadata about current posts
type PostCounter struct {
	Count int    `json:"Count"`
	List  []Meta `json:"List"`
}
