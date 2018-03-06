package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func editPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	editid, err := strconv.Atoi(vars["editid"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	exsistingpost, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	newpost := postdata{}
	err = json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	if exsistingpost.EditID != editid {
		fmt.Fprintln(w, "Edit id did not match actual Edit id")
		return
	}
	newpost.ID = exsistingpost.ID
	newpost.EditID = exsistingpost.EditID
	err = saveChanges(newpost)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(newpost)
}
func deletePostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	exsistingpost, err := readpostDB(ID)
	if exsistingpost.EditID != editid {
		fmt.Fprintf(w, "Edit ID %v is not correct", editid)
		return
	}
	err = deletepost(exsistingpost)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "No such post")
	}
}
func postcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	json.NewEncoder(w).Encode(posts)
}
func requestPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	result, err := readpostDB(i)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	err = incrementViewCounter(result.ID)
	if err != nil {
		log.Println(err)
	}
}
func createPostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := postdata{}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = genFromSeed()
	newpost.EditID = genFromSeed()
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "No data posted!")
		return
	}
	defer r.Body.Close()
	createPostDB(newpost)

	json.NewEncoder(w).Encode(newpost)
	r.Body.Close()
}
