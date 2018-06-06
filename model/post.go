package model

import (
	"html/template"
	"time"
)

type Post struct {
	ID int `json:"ID" gorm:"Column:id;primary_key"`
	Data
	Meta
	Edit
	Crypto
}

//PostData contains all the info related to the posts
type Data struct {
	Title     string `json:"Title" gorm:"Column:title"`
	PostPerms int    `json:"PostPerms,string" gorm:"Column:postperms"`
	Content   string `json:"Content" gorm:"Column:text"`
}

type Crypto struct {
	Hash     string `json:"Password" gorm:"Column:hash"`
	Salt     string `json:"Salt" gorm:"Column:salt"`
	AuthHash string `json:"authHash" gorm:"-"`
	Key      string `json:"Key" gorm:"Column:key"`
}

type Meta struct {
	Views      int        `json:"Views" gorm:"Column:views"`
	CreateTime time.Time  `json:"Time" gorm:"Column:created_at"`
	UpdateTime time.Time  `json:"Updated" gorm:"Column:updated_at"`
	DeleteTime *time.Time `json:"Deleted" gorm:"Column:deleted_at"`
}

type Markdown struct {
	Md      template.HTML `json:"" gorm:"-"`
	TitleMD template.HTML `json:"" gorm:"-"`
}

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

//PostCreate will be used as a struct when posts are created, as to try to reduce the use of the god struct PostData
type PostCreate struct {
	csrfField string
}

//PostDecrypt is used for decrypting post content
type PostDecrypt struct {
	Post
	Mode  string
	Token string
}

//PostCounter is used on index to provide some metadata about current posts
type PostCounter struct {
	Count int    `json:"Count"`
	Meta  []Data `json:"Meta"`
}
