package main

import (
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
	db := establishConn()
	err := db.QueryRow("SELECT ID FROM TEXT LIMIT 1;").Scan(&post)
	if err != nil {
		return 0, err
	}
	db.Close()
	return post, err
}
