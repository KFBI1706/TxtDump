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
	router.HandleFunc("/", logging(displayIndex)).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", logging(postcounterAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(requestPostAPI)).Methods("GET")
	router.HandleFunc("/post/{id}/request", logging(requestPostHTML)).Methods("GET")
	router.HandleFunc("/api/v1/post/create", logging(createPostAPI)).Methods("POST")
	router.HandleFunc("/post/create", logging(createPostWeb))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	err = http.ListenAndServe(":1337", router)
	if err != nil {
		log.Println(err)
	}
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
