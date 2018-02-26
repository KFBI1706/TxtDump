package main

import (
	"math/rand"
)

//Probably not the best way to do this
func genFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	return num
}
func findpostfortest() (int, error) {
	var post int
	db := establishConn()
	err := db.QueryRow("SELECT pubid FROM text WHERE id = $1", 1).Scan(&post)
	if err != nil {
		return 0, err
	}
	db.Close()
	return post, err
}
