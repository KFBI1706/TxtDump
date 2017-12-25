package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	testDBConnection()
	router := mux.NewRouter()
	router.HandleFunc("/test", routerTest)
	router.HandleFunc("/post/{id}/request", requestPost).Methods("GET")
	router.HandleFunc("/post/create", createPost).Methods("POST")
	router.HandleFunc("/random/test", requestPostID).Methods("GET")
	http.ListenAndServe(":1337", router)
}
