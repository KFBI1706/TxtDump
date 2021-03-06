package sql

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/KFBI1706/TxtDump/config"
	"github.com/KFBI1706/TxtDump/model"
	_ "github.com/lib/pq"
)

var posts []model.Post

func generateRandomPostData() {
	for len(posts) < 10 {
		rand.Seed(time.Now().UTC().UnixNano())
		post := model.Post{Data: model.Data{Content: randomString(20), PostPerms: 2}, Meta: model.Meta{Views: 0}}
		post.Data.Title = randomString(10)
		post.ID = rand.Intn(9999999-1000000) + 1000000
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

func TestMain(m *testing.M) {
	conf := config.ParseConfig("development")
	config.InitDB(conf.DBStringLocation)
	code := m.Run()
	os.Exit(code)
}

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

func TestCreateReadSaveDelete(t *testing.T) {
	generateRandomPostData()
	for _, post := range posts {
		if !CheckForDuplicateID(post.ID) {
			t.Error("Duplicate ID")
		}
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
		readpost.Data.Content = randomString(20)
		err = SaveChanges(readpost)
		if err != nil {
			t.Error(err)
		}
		savedpost, err := ReadPostDB(readpost.ID)
		if err != nil || savedpost.Data.Content != readpost.Data.Content {
			t.Error("something went wrong saving changes")
		}
		err = DeletePost(readpost)
		if err != nil {
			t.Error(err)
		}
	}
}
