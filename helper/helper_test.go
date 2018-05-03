package helper_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/KFBI1706/Txtdump/helper"
)

func TestIDGenerator(t *testing.T) {
	var ids []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		id := helper.GenFromSeed()
		ids = append(ids, id)
	}
}
