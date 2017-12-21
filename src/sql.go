package main

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
)

func readDBstring(file string) {
	ioutil.ReadFile(file)
}
func establishConn() {

	db, err := sql.Open("postgres")
}
