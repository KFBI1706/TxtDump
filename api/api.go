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

func EditPostAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := html.ProcessRequest(w, r)
	newpost := model.PostData{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID

	if valid := crypto.CheckPass(existingpost.Hash, existingpost.ID, existingpost.PostPerms); valid {
		err = sql.SaveChanges(newpost)
	}
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Wrong Password")
		return
	}
	json.NewEncoder(w).Encode(newpost)
}

func DeletePostAPI(w http.ResponseWriter, r *http.Request) {
	post := html.ProcessRequest(w, r)
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
	}

	if valid := crypto.CheckPass(post.Hash, post.ID, post.PostPerms); valid {
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

func PostcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := model.PostCounter{Count: sql.CountPosts()}
	posts, err := sql.PostMetas()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(posts)
}

func RequestPostAPI(w http.ResponseWriter, r *http.Request) {
	result := html.ProcessRequest(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if result.PostPerms == 3 {
		fmt.Fprintf(w, "This post is password protected mah dude you need to POST the Hash")
		return
	}
	result.Hash = ""
	result.Salt = ""
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	err = sql.IncrementViewCounter(result.ID)
	if err != nil {
		log.Println(err)
	}
}

func RequestPostWithPassAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := html.ProcessRequest(w, r)
	newpost := model.PostData{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID
	json.NewEncoder(w).Encode(newpost)
}

func CreatePostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := model.PostData{}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = helper.GenFromSeed()
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Something went wrong")
		return
	}
	if newpost.PostPerms == 0 || newpost.PostPerms > 3 {
		newpost.PostPerms = 2
	}
	crypto.SecurePost(&newpost, newpost.Hash)
	sql.CreatePostDB(newpost)
	json.NewEncoder(w).Encode(newpost)
}
