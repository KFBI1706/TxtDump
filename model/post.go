package model

import (
	"html/template"
	"time"
)

type PostData struct {
	ID        int           `json:"ID"`
	EditID    int           `json:"EditID"`
	Hash      string        `json:"Password"`
	Salt      string        `json:"Salt"`
	AuthHash  string        `json:"authHash"`
	Key       string        `json:"Key"`
	PostPerms int           `json:"PostPerms,string"`
	Content   string        `json:"Content"`
	Md        template.HTML `json:""`
	Title     string        `json:"Title"`
	TitleMD   template.HTML `json:""`
	Time      time.Time     `json:"Time"`
	Views     int           `json:"Views"`
}

type PostDecrypt struct {
	ID   int
	Mode string
}
