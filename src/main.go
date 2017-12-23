package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	establishConn()
	router := mux.NewRouter()
	router.HandleFunc("/test", routerTest)
	router.HandleFunc("/post/{id}/request", requestPost)
	router.HandleFunc("/post/{id}/create", createPost)
	router.HandleFunc("/random/test", requestPostID)
	http.ListenAndServe(":1337", router)
}
