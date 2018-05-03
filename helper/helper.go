package helper

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

//Probably not the best way to do this
func genFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	for !checkForDuplicateID(num) {
		num = rand.Intn(9999999-1000000) + 1000000
	}
	return num
}

func HexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
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
	fmt.Println(res)
	if err != nil {
		return err
	}
	db.Close()
	return nil
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
	return nil
}

func determinePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
