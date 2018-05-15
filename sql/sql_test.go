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

func TestGenerateRandomPostData(t *testing.T) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		post := model.PostData{}
		post.ID = rand.Intn(9999999-1000000) + 1000000
		posts = append(posts, post)
	}
	fmt.Println(posts)
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

func TestReadPostDB(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    model.PostData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadPostDB(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPostDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadPostDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveChanges(t *testing.T) {
	type args struct {
		post model.PostData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveChanges(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("SaveChanges() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountPosts(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPosts(); got != tt.want {
				t.Errorf("CountPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostMetas(t *testing.T) {
	tests := []struct {
		name    string
		want    model.PostCounter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostMetas()
			if (err != nil) != tt.wantErr {
				t.Errorf("PostMetas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostMetas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	type args struct {
		post model.PostData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIncrementViewCounter(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncrementViewCounter(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("IncrementViewCounter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckForDuplicateID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckForDuplicateID(tt.args.id); got != tt.want {
				t.Errorf("CheckForDuplicateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProp(t *testing.T) {
	type args struct {
		prop string
		id   int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProp(tt.args.prop, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProp() = %v, want %v", got, tt.want)
			}
		})
	}
}
