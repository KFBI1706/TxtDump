package sql

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var posts []model.PostData

func generateRandomPostData() {
	for len(posts) < 10 {
		rand.Seed(time.Now().UTC().UnixNano())
		post := model.PostData{Title: randomString(10), Content: randomString(20), Hash: randomString(8), Key: randomString(8), Salt: randomString(8), PostPerms: 2, Views: 0}
		post.ID = rand.Intn(9999999-1000000) + 1000000
		post.EditID = rand.Intn(9999999-1000000) + 1000000
		posts = append(posts, post)

	}
}
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))

	}
	return string(bytes)
}

func TestReadDBstring(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Totally not here for code coverage", args: args{filename: "testdb/test.sql"}, want: "Testerino", wantErr: false},
		{name: "Test on non exsisting file", args: args{filename: "testdb/xd.sql"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDBstring(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDBstring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadDBstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

//This will fail if there is not a dbstring file in the SQL folder. This should probably be fixed by config file
func TestDBConnectionfunc(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Run once", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TestDBConnection(); (err != nil) != tt.wantErr {
				t.Errorf("TestDBConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEstablishConn(t *testing.T) {
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
	}{
		//	{name: "Establish conn to prod", want: gorm.DB{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EstablishConn()
			defer got.Close()
			if (err != nil) != tt.wantErr {
				t.Errorf("EstablishConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EstablishConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateReadSaveDelete(t *testing.T) {
	generateRandomPostData()
	for _, post := range posts {
		err := CreatePostDB(post)
		if err != nil {
			t.Error(err)
		}
		err = IncrementViewCounter(post.ID)
		if err != nil {
			t.Error(err)
		}
		readpost, err := ReadPostDB(post.ID)
		if err != nil {
			t.Error(err)
		}
		readpost.Content = randomString(20)
		err = SaveChanges(readpost)
		if err != nil {
			t.Error(err)
		}
		err = DeletePost(readpost)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(readpost)
	}
}
