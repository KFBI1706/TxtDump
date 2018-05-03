package helper

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/KFBI1706/Txtdump/sql"
)

//Probably not the best way to do this
func GenFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	for !sql.CheckForDuplicateID(num) {
		num = rand.Intn(9999999-1000000) + 1000000
	}
	return num
}

func findpostfortest() (int, error) {
	var post int
	db, err := sql.EstablishConn()
	err = db.QueryRow("SELECT ID FROM TEXT LIMIT 1;").Scan(&post)
	if err != nil {
		return 0, err
	}
	db.Close()
	return post, err
}
func SetupDB() error {
	db, err := sql.EstablishConn()
	if err != nil {
		return err
	}
	sql, err := sql.ReadDBstring("sql/db.sql")
	if err != nil {
		return err
	}
	res, err := db.Exec(sql)
	fmt.Println(res)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func ClearOutDB() error {
	db, err := sql.EstablishConn()
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE text;")
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func DeterminePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
