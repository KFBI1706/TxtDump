package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestIDGenerator(t *testing.T) {
	var ids []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		id := genFromSeed()
		ids = append(ids, id)
	}
}
