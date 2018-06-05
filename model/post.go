package model

import (
	"html/template"
	"time"
)

//PostData contains all the info related to the posts
type PostData struct {
	PostEdit
	PostMeta
	PostRender
	ID        int    `json:"ID" gorm:"Column:id;primary_key"`
	EditID    int    `json:"EditID" gorm:"Column:editid"`
	Hash      string `json:"Password" gorm:"Column:hash"`
	Salt      string `json:"Salt" gorm:"Column:salt"`
	AuthHash  string `json:"authHash" gorm:"-"`
	Key       string `json:"Key" gorm:"Column:key"`
	PostPerms int    `json:"PostPerms,string" gorm:"Column:postperms"`
	Content   string `json:"Content" gorm:"Column:text"`
}
type PostEdit struct {
	Title string `json:"Title" gorm:"Column:title"`
}
type PostMeta struct {
	CreateTime time.Time  `json:"Time" gorm:"Column:created_at"`
	UpdateTime time.Time  `json:"Updated" gorm:"Column:updated_at"`
	DeleteTime *time.Time `json:"Deleted" gorm:"Column:deleted_at"`
	Views      int        `json:"Views" gorm:"Column:views"`
}
type PostRender struct {
	MD      template.HTML `json:"" gorm:"-"`
	TitleMD template.HTML `json:"" gorm:"-"`
}

//PostCreate will be used as a struct when posts are created, as to try to reduce the use of the god struct PostData
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
	Meta  []PostData `json:"Meta"`
}
