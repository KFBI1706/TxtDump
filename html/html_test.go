package html

import (
	"os"
	"testing"

	"github.com/KFBI1706/TxtDump/config"
)

//Production tests for ssl
func TestMain(m *testing.M) {
	conf := config.ParseConfig("development")
	config.InitDB(conf.DBStringLocation)
	code := m.Run()
	os.Exit(code)
}
