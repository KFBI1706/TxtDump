package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func editPostAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := processRequest(w, r)
	newpost := postData{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID
	err = saveChanges(newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Wrong Password")
		return
	}
	json.NewEncoder(w).Encode(newpost)
}

func deletePostAPI(w http.ResponseWriter, r *http.Request) {
	post := processRequest(w, r)
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
	}
	err = deletepost(post)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Either the post does not exist or the password is wrong")
		return
	}
	fmt.Fprintf(w, "Post %v successfully deleted", post.ID)
}

func postcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	posts, err := postMeta()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(posts)
}

func requestPostAPI(w http.ResponseWriter, r *http.Request) {
	result := processRequest(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if result.PostPerms == 3 {
		fmt.Fprintf(w, "This post is password protected mah dude you need to POST the Hash")
		return
	}
	result.EditID = 0
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	err = incrementViewCounter(result.ID)
	if err != nil {
		log.Println(err)
	}
}

func requestPostWithPassAPI(w http.ResponseWriter, r *http.Request) {
	existingpost := processRequest(w, r)
	newpost := postData{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	newpost.ID = existingpost.ID
	json.NewEncoder(w).Encode(newpost)
}

func createPostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := postData{}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = genFromSeed()
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Something went wrong")
		return
	}
	securePost(&newpost, newpost.Hash)
	createPostDB(newpost)
	json.NewEncoder(w).Encode(newpost)
}
