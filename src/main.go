package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", routerTest)
	router.HandleFunc("/request/post/{id}", requestPost)
	router.HandleFunc("/random/test", requestPostID)
	http.ListenAndServe(":1337", router)

}
