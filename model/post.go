package model

import (
	"html/template"
	"time"
)

//PostData contains all the info related to the posts
type PostData struct {
	ID         int           `json:"ID" gorm:"Column:id;primary_key"`
	EditID     int           `json:"EditID" gorm:"Column:editid"`
	Hash       string        `json:"Password" gorm:"Column:hash"`
	Salt       string        `json:"Salt" gorm:"Column:salt"`
	AuthHash   string        `json:"authHash" gorm:"-"`
	Key        string        `json:"Key" gorm:"Column:key"`
	PostPerms  int           `json:"PostPerms,string" gorm:"Column:postperms"`
	Content    string        `json:"Content" gorm:"Column:text"`
	Md         template.HTML `json:"" gorm:"-"`
	Title      string        `json:"Title" gorm:"Column:title"`
	TitleMD    template.HTML `json:"" gorm:"-"`
	CreateTime time.Time     `json:"Time" gorm:"Column:created_at"`
	UpdateTime time.Time     `json:"Updated" gorm:"Column:updated_at"`
	DeleteTime *time.Time    `json:"Deleted" gorm:"Column:deleted_at"`
	Views      int           `json:"Views" gorm:"Column:views"`
}
type PostCreate struct {
	csrfField string
}

//PostDecrypt is used for decrypting post content
type PostDecrypt struct {
	ID    int
	Mode  string
	Token string
}

//PostCounter is used on index to provide some metadata about current posts
type PostCounter struct {
	Count int        `json:"Count"`
	Meta  []PostMeta `json:"Meta"`
}

//PostMeta is used as array for Post Metadata in PostCoutner
type PostMeta struct {
	PostID    int    `json:"ID"`
	Title     string `json:"Title"`
	Views     int    `json:"View"`
	PostPerms int    `json:"PostPerms"`
}
