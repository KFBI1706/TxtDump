package helper_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/KFBI1706/TxtDump/config"
	"github.com/KFBI1706/TxtDump/helper"
)

func TestMain(m *testing.M) {
	conf := config.ParseConfig("development")
	config.InitDB(conf.DBStringLocation)
	code := m.Run()
	os.Exit(code)
}

func TestIDGenerator(t *testing.T) {
	var ids []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		id := helper.GenFromSeed()
		ids = append(ids, id)
	}
}
