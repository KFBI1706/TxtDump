package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/KFBI1706/TxtDump/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	TestString        string
	TestConfiguration model.Configuration
	CorruptedFile     string
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestMain(m *testing.M) {
	TestString = "hello"
	err := ioutil.WriteFile("/tmp/testfile", []byte(TestString), 0644)
	if err != nil {
		panic(err)
	}
	pth, _ := findConfig("_corrupted")
	err = ioutil.WriteFile(pth, []byte("{"), 0644)
	if err != nil {
		panic(err)
	}
	retCode := m.Run()
	err = os.Remove("/tmp/testfile")
	if err != nil {
		panic(err)
	}
	err = os.Remove(pth)
	if err != nil {
		panic(err)
	}
	os.Exit(retCode)
}

func Test_readFile(t *testing.T) {
	type args struct {
		pth string
	}
	tests := []struct {
		name        string
		args        args
		wantContent string
	}{
		{"Undefined path", args{pth: "lel"}, ""},
		{"Defined path", args{pth: "/tmp/testfile"}, TestString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotContent := readFile(tt.args.pth); gotContent != tt.wantContent {
				t.Errorf("readFile() = '%v', want '%v'", gotContent, tt.wantContent)
			}
		})
	}
}

func TestInitDB(t *testing.T) {
	type args struct {
		dbstringLocation string
	}
	tests := []struct {
		name      string
		path      string
		wantPanic bool
	}{
		{"Correct file-path", ParseConfig("development").DBStringLocation, false}, {"Wrong file-path", "lol", true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanic(t, func() { InitDB(tt.path) })
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	type args struct {
		env string
	}
	//root := projectRoot()
	tests := []struct {
		name       string
		args       args
		wantConfig model.Configuration
		wantPanic  bool
	}{
		{"Testing missing file", args{env: "asdf"}, model.Configuration{}, true},

		{"Testing corrupted file", args{env: "_corrupted"}, model.Configuration{}, true},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanic(t, func() { ParseConfig(tt.args.env) })
			} else if gotConfig := ParseConfig(tt.args.env); !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("ParseConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
