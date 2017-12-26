package main

import "testing"

func testdbconn(t *testing.T) {
	err := testDBConnection()
	if err != nil {
		t.Error(err)
	}
	redpost := readpostDB(1)
	if redpost.Content == "" {
		t.Error("Something went wrong reading from DB")
	}
}
