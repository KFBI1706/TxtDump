package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type postresp struct {
	ID      int
	Content string
	Sucsess bool //Optional
}

func routerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jallo")
}

//Probably not the safest way of doing this but works for now
func requestPostID(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	generatedID := genFromSeed()
	requestedid := postresp{Content: "Your ID", ID: generatedID}
	json.NewEncoder(w).Encode(requestedid)
}
func requestPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	post := postresp{ID: i, Content: "not implemented yet", Sucsess: true}
	json.NewEncoder(w).Encode(post)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	newpost := postresp{}
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		fmt.Fprint(w, "No data posted!")
		return
	}
	newpost.Sucsess = true
	defer r.Body.Close()
	json.NewEncoder(w).Encode(newpost)
}
