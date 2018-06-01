package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DB : is a exported variable for access from other packages, as this is the package that inits the DB
var DB *gorm.DB

//InitDB function reads the file location contents, which should contain database parameters and uses them to create a pointer to DB
//takes dbstringLocation (string) as a parameter, which shuold be a full path to the file
//nothing is returned directly, but if the function executes correctly the exported DB variable should be a *DB
func checkFile(pth string) error {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return err
	}
	return nil
}

func readFile(pth string) string {
	if checkFile(pth) != nil {
		return ""
	}
	bytes, _ := ioutil.ReadFile(pth)
	return string(bytes)
}

func InitDB(dbstringLocation string) {
	var err error
	if DB, err = gorm.Open("postgres", readFile(dbstringLocation)); err != nil {
		log.Panic(err)
	}
	var postData model.PostData
	if !DB.HasTable(&postData) {
		DB.AutoMigrate(&postData)
	}
}

func projectRoot() string {
	gp := os.Getenv("GOPATH")
	pth := path.Join(gp, "src/github.com/KFBI1706/TxtDump")
	return pth
}
func findConfig(env string) (string, error) {
	pth := projectRoot()
	return path.Join(pth, "config."+env+".json"), nil
}

func findDBString(path string) string {
	return path + "/dbstring"
}

//ParseConfig does all the ParseConfig magic, with eventual support for environment variables etc.
//takes env (string) as an argument which is used to denote the environment the application is in
//returns model.Configuration struct containing various information about used by the aplication, see model/conf.go
func ParseConfig(env string) (config model.Configuration) {
	var file *os.File
	conf, err := findConfig(env)
	file, err = os.Open(conf)
	if err != nil {
		panic(err)

	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}
	config.Path = projectRoot()
	config.DBStringLocation = findDBString(config.Path)
	return
}
