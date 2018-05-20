package helper

import (
	"math/rand"
	"strconv"

	"github.com/KFBI1706/TxtDump/sql"
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

/*DeterminePerms converts the postPerm string to int
one input argument string (postperm)
two return arguments which returns the int (num), and error (err)*/
func DeterminePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
