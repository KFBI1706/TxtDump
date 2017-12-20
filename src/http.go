package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type postreq struct {
	ID int
}
type postresp struct {
	ID      int
	Content string
}

func routerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Im working maayn")
}

//Probably not the safest way of doing this but works for now
func requestPostID(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	generatedID := genFromSeed()
	fmt.Fprintf(w, "Seed is: %v", generatedID)
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
	post := postreq{ID: i}
	fmt.Fprintf(w, "U requested id: %v", post.ID)
}
