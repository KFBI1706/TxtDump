package main

import (
	"fmt"
	"log"
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
	fmt.Printf("The first post in the DB has the id: %v\n", post)
	redpost := readpostDB(post)
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
func TestPostCreateEditDelete(t *testing.T) {
	post := postdata{Content: "Post Generated for testing", Title: "Post Generated for testing", ID: 1000, EditID: 1000, Time: time.Now()}
	createPostDB(post)
	post.Content = "Second Phase"
	err := saveChanges(post)
	if err != nil {
		t.Error(err)
	}
	err = incrementViewCounter(post.ID)
	if err != nil {
		log.Println(err)
	}
	err = deletepost(post)
	if err != nil {
		t.Error(err)
	}
}
