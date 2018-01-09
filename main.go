package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := testDBConnection()
	posts := countPosts()
	log.Printf("%v Post(s) Currently in DB\n", posts)
	if err != nil {
		log.Println(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", displayIndex).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", postcounterAPI).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", requestPostAPI).Methods("GET")
	router.HandleFunc("/post/{id}/request", requestPostHTML).Methods("GET")
	router.HandleFunc("/api/v1/post/create", createPost).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	err = http.ListenAndServe(":1337", router)
	if err != nil {
		log.Println(err)
	}
}
