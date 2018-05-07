package helper

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/KFBI1706/Txtdump/sql"
)

/*GenFromSeed generates a guaranteed random number that's not already in the database
no input arguments
one return argument int (num)*/
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

/*SetupDB is used to setup the database
no input arguments
returns error*/
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

/*SetupTestDB is used to setup the database
no input arguments
returns error*/
func SetupTestDB() error {
	db, err := sql.EstablishConn()
	if err != nil {
		return err
	}
	sql, err := sql.ReadDBstring("sql/testdb/test.sql")
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

/*ClearOutDB is used to clear a table
no input arguments
returns error*/
func ClearOutDB() error {
	db, err := sql.EstablishConn()
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE text, text_test;")
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

/*DeterminePerms converts the postPerm string to int
one input argument string (postperm)
two return arugments which returns the int (num), and error (err)*/
func DeterminePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
