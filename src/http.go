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
	post := postresp{ID: i, Content: "not implemented yet"}
	json.NewEncoder(w).Encode(post)
}
