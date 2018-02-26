package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDBconn(t *testing.T) {
	err := testDBConnection()
	if err != nil {
		t.Error(err)
	}
	post, err := findpostfortest()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("The first post in the DB has the pubid: %v\n", post)
	redpost := readpostDB(post)
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
func TestPostCreateEditDelete(t *testing.T) {
	post := postdata{Content: "Post Generated for testing", Title: "Post Generated for testing", PubID: 1000, EditID: 1000, Time: time.Now()}
	createPostDB(post)
	post.Content = "Second Phase"
	err := saveChanges(post)
	if err != nil {
		t.Error(err)
	}
	err = deletepost(post)
	if err != nil {
		t.Error(err)
	}
}
