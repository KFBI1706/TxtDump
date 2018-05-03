package sql_test

import (
	"fmt"
	"testing"
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

//func TestPostCreateEditDelete(t *testing.T) {
//	newpost := postData{Title: "Test Post", Content: "Test Content", PostPerms: 2, Hash: "Testpass"}
//	rand.Seed(time.Now().UnixNano())
//	newpost.ID = genFromSeed()
//	securePost(&newpost, newpost.Hash)
//	err := createPostDB(newpost)
//	if err != nil {
//		t.Error(err)
//	}
//	createdpost, err := readpostDB(newpost.ID)
//	createdpost.Hash = "Testpass"
//	if err != nil {
//		t.Error(err)
//	}
//	if createdpost.Title != newpost.Title {
//		t.Error(createdpost.Content, newpost.Title)
//	}
//	err = deletepost(createdpost)
//	if err != nil {
//		t.Error(err)
//	}
//}
//
