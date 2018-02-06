package main

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
	fmt.Println(post)
	redpost := readpostDB(post)
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
