package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := testDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v Post(s) Currently in DB\n", countPosts())
	router := mux.NewRouter()
	router.HandleFunc("/", logging(displayIndex)).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", logging(postcounterAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(requestPostAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/create", logging(createPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/edit/{editid}", logging(editPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/delete/{editid}", logging(deletePostAPI))
	router.HandleFunc("/post/{id}/request", logging(requestPostWeb)).Methods("GET")
	router.HandleFunc("/post/{id}/edit/{editid}", logging(editPost))
	router.HandleFunc("/post/{id}/edit/{editid}/post", logging(edit))
	router.HandleFunc("/post/{id}/delete/{editid}/", logging(deletePostWeb))
	router.HandleFunc("/post/create", logging(createPostTemplateWeb))
	router.HandleFunc("/post/create/new", logging(createPostWeb))
	router.HandleFunc("/documentation", logging(documentation))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	router.Walk(routerWalk)
	err = http.ListenAndServe(":1337", router)
	if err != nil {
		log.Println(err)
	}
}
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.URL.Path)
		f(w, r)
	}
}
