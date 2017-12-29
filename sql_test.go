package main

import "testing"

func TestDBconn(t *testing.T) {
	err := testDBConnection()
	if err != nil {
		t.Error(err)
	}
	redpost := readpostDB(9175728)
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
