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

//DB : is a exported variable for access from other packages, as this is the package that inits the DB
var DB *gorm.DB

//InitDB function reads the file location contents, which should contain database parameters and uses them to create a pointer to DB
//takes dbstringLocation (string) as a parameter, which shuold be a full path to the file
//nothing is returned directly, but if the function executes correctly the exported DB variable should be a *DB
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

//ParseConfig does all the ParseConfig magic, with eventual support for environment variables etc.
//takes env (string) as an argument which is used to denote the environment the application is in
//returns model.Configuration struct containing various information about used by the aplication, see model/conf.go
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
