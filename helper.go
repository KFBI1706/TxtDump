package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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
func securePass(ps string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}
func checkPass(ps string, id int) error {
	err := bcrypt.CompareHashAndPassword(getHashedPS(id), []byte(ps))
	return err
}
func clearOutDB() error {
	db, err := establishConn()
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE text")
	if err != nil {
		return err
	}
	return err
}
func determinePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
