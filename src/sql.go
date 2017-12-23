package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

//Since the Postgresql Go libary just uses a string for info i just read a file with the private database info in it as a string with this see readme.md for more
func readDBstring(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	returnstring := string(file)

	return returnstring
}

func establishConn() {
	dbstring := readDBstring("dbstring")
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connetcion sucsessfully established")
}
