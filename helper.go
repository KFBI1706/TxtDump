package main

import (
	"fmt"
	"log"
	"math/rand"
)

//Probably not the best way to do this
func genFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	if checkForDuplicateID(num) {
		return num
	}
	num = rand.Intn(9999999-1000000) + 1000000
	log.Println(num)
	return num
}
func findpostfortest() (int, error) {
	var post int
	db, err := establishConn()
	err = db.QueryRow("SELECT ID FROM TEXT LIMIT 1;").Scan(&post)
	if err != nil {
		return 0, err
	}
	db.Close()
	return post, err
}
func setupDB() error {
	db, err := establishConn()
	if err != nil {
		return err
	}
	sql, err := readDBstring("sql/db.sql")
	if err != nil {
		return err
	}
	res, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(res)
	db.Close()
	return nil
}
