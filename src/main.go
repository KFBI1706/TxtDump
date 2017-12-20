package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", routerTest)
	router.HandleFunc("/request/post/{id}", requestPost)
	http.ListenAndServe(":1337", router)
}
