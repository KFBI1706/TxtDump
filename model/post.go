package model

import (
	"html/template"
	"time"
)

//Post is the "master" struct
type Post struct {
	ID int `json:"ID" gorm:"Column:id;primary_key"`
	Data
	Meta
	Edit
	Crypto
}

//Data contains all the info related to the posts
type Data struct {
	Title     string `json:"Title" gorm:"Column:title"`
	PostPerms int    `json:"PostPerms,string" gorm:"Column:postperms"`
	Content   string `json:"Content" gorm:"Column:text"`
}

//Crypto contains the cryptographic aspect of the post
type Crypto struct {
	Hash     string `json:"Password" gorm:"Column:hash"`
	Salt     string `json:"Salt" gorm:"Column:salt"`
	AuthHash string `json:"authHash" gorm:"-"`
	Key      string `json:"Key" gorm:"Column:key"`
}

//Meta Contains most of the metadata views timestamps etc
type Meta struct {
	Views      int        `json:"Views" gorm:"Column:views"`
	CreateTime time.Time  `json:"Time" gorm:"Column:created_at"`
	UpdateTime time.Time  `json:"Updated" gorm:"Column:updated_at"`
	DeleteTime *time.Time `json:"Deleted" gorm:"Column:deleted_at"`
}

//Markdown Contains the data used for rendering the post on the HTML frontend
type Markdown struct {
	MD      template.HTML `json:"" gorm:"-"`
	TitleMD template.HTML `json:"" gorm:"-"`
}

//PostNew is used for registering a new post
type PostNew struct {
	Post
	//ID int `json:"ID" gorm:"Column:id;primary_key"`
	//edit PostEdit
	//meta PostMeta
	//data PostData
}

type Edit struct {
	EditID int `json:"EditID" gorm:"Column:editid"`
}

//PostCounter is used on index to provide some metadata about current posts
type PostCounter struct {
	Count int    `json:"Count"`
	Meta  []Data `json:"Meta"`
}
