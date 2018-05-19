package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB(dbstringLocation string) {
	dbstring, err := ioutil.ReadFile(dbstringLocation)
	if err != nil {
		log.Panic(err)
	}
	if DB, err = gorm.Open("postgres", string(dbstring)); err != nil {
		log.Panic(err)
	}
}

func findConfig(env string) string {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/KFBI1706/TxtDump")
	return path.Join(ap, "config."+env+".json")
}

//does all the ParseConfig magic, with eventual support for environment variables etc.
func ParseConfig(env string) (config model.Configuration) {
	file, err := os.Open(findConfig(env))
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		panic(err)
	}
	config.DBStringLocation = config.Path + "/dbstring"
	return
}
