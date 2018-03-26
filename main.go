package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := testDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	dbsetup := flag.Bool("setupdb", false, "Setup db when running")
	port := flag.Int("port", 1337, "for using a custom port")
	flag.Parse()
	addr := fmt.Sprintf(":%v", *port)
	if *dbsetup == true {
		err = setupDB()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%v Post(s) Currently in DB\n", countPosts())
	router := mux.NewRouter()
	router.HandleFunc("/", logging(displayIndex)).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", logging(postcounterAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(requestPostAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/create", logging(createPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/edit", logging(editPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/delete", logging(deletePostAPI))
	router.HandleFunc("/post/{id}/request", logging(requestPostWeb)).Methods("GET")
	router.HandleFunc("/post/{id}/edit", logging(editPost))
	router.HandleFunc("/post/{id}/edit/post", logging(edit)).Methods("POST")
	router.HandleFunc("/post/{id}/delete", logging(deletePostWeb)).Methods()
	router.HandleFunc("/post/create", logging(createPostTemplateWeb))
	router.HandleFunc("/post/create/new", logging(createPostWeb))
	router.HandleFunc("/documentation", logging(documentation))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	router.Walk(routerWalk)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.URL.Path)
		f(w, r)
	}
}
func routerWalk(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	t, err := route.GetPathTemplate()
	if err != nil {
		return err
	}
	fmt.Println(t)
	return nil
}
