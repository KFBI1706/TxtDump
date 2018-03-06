package main

import (
	"fmt"
	"log"
	"math/rand"
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
	redpost, err := readpostDB(post)
	if err != nil {
		t.Error(err)
	}
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
func TestPostCreateEditDelete(t *testing.T) {
	post := postdata{Content: "Post Generated for testing", Title: "Post Generated for testing", EditID: 1000, Time: time.Now()}
	rand.Seed(time.Now().UnixNano())
	post.ID = genFromSeed()
	fmt.Println(post.ID)
	if checkForDuplicateID(post.ID) != true {
		t.Errorf("Post with ID %v Already exsits", post.ID)
	}
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
