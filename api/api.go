package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/KFBI1706/TxtDump/crypto"
	"github.com/KFBI1706/TxtDump/helper"
	"github.com/KFBI1706/TxtDump/html"
	"github.com/KFBI1706/TxtDump/model"
	"github.com/KFBI1706/TxtDump/sql"
)

//EditPostAPI is the HTTP handler for editing posts
func EditPostAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := html.ProcessRequest(w, r)
	newpost := model.Post{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID

	if valid := crypto.CheckPass(existingpost.Crypto.Hash, existingpost.ID, existingpost.Data.PostPerms); valid {
		if err = sql.SaveChanges(newpost); err != nil {
			panic(err)
		}
	} else {
		log.Println(err)
		fmt.Fprint(w, "Wrong Password")
		return
	}
	json.NewEncoder(w).Encode(newpost)
}

//DeletePostAPI is the HTTP handler for deleteing posts
func DeletePostAPI(w http.ResponseWriter, r *http.Request) {
	post := html.ProcessRequest(w, r)
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
	}

	if valid := crypto.CheckPass(post.Crypto.Hash, post.ID, post.Data.PostPerms); valid {
		err = sql.DeletePost(post)
	}
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Either the post does not exist or the password is wrong")
		return
	}
	fmt.Fprintf(w, "Post %v successfully deleted", post.ID)
	return
}

//PostcounterAPI is returns some metadata about data in the DB
func PostcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := model.PostCounter{Count: sql.CountPosts()}
	//	//posts, err := sql.PostMetas()
	//	if err != nil {
	//		log.Println(err)
	//	}
	json.NewEncoder(w).Encode(posts)
}

//RequestPostAPI is the HTTP handler for "reading" posts
func RequestPostAPI(w http.ResponseWriter, r *http.Request) {
	result := html.ProcessRequest(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if result.Data.PostPerms == 3 {
		fmt.Fprintf(w, "This post is password protected mah dude you need to POST the Hash")
		return
	}
	result.Crypto.Hash = ""
	result.Crypto.Salt = ""
	defer sql.IncrementViewCounter(result.ID)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

//RequestPostWithPassAPI is for reading posts that are password encrypted
func RequestPostWithPassAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := html.ProcessRequest(w, r)
	newpost := model.Post{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID
	json.NewEncoder(w).Encode(newpost)
}

//CreatePostAPI is the HTTP handler for registering posts via the API
func CreatePostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := model.Post{}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = helper.GenFromSeed()
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Something went wrong")
		return
	}
	if newpost.Data.PostPerms == 0 || newpost.Data.PostPerms > 3 {
		newpost.Data.PostPerms = 2
	}
	crypto.SecurePost(&newpost, newpost.Crypto.Hash)
	sql.CreatePostDB(newpost)
	json.NewEncoder(w).Encode(newpost)
}
