package main

import (
	"math/rand"
)

//Probably not the best way to do this
func genFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	return num
}
