package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	err := testDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	dbdrop := flag.Bool("dropdb", false, "Drop current table and all data")
	dbsetup := flag.Bool("setupdb", false, "Setup db when running")
	port := flag.Int("port", 1337, "for using a custom port")
	flag.Parse()
	addr := fmt.Sprintf(":%v", *port)
	if *dbdrop || *dbsetup {
		if *dbdrop {
			err = clearOutDB()
			if err != nil {
				log.Fatal(err)
			}
		}
		if *dbsetup {
			err = setupDB()
			if err != nil {
				log.Println(err)
			}
		}
		os.Exit(3)
	}
	log.Printf("%v Post(s) Currently in DB\n", countPosts())
	router := mux.NewRouter()
	router.HandleFunc("/", logging(displayIndex)).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", logging(postcounterAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(requestPostAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(requestPostWithPassAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/create", logging(createPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/edit", logging(editPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/delete", logging(deletePostAPI))
	router.HandleFunc("/post/{id}/request", logging(requestPostWeb)).Methods("GET")
	router.HandleFunc("/post/{id}/request/decrypt", logging(requestPostDecrypt)).Methods("POST")
	router.HandleFunc("/post/{id}/edit", logging(editPostTemplate))
	router.HandleFunc("/post/{id}/edit/decrypt", logging(editPostDecrypt))
	router.HandleFunc("/post/{id}/edit/post", logging(editPostForm)).Methods("POST")
	router.HandleFunc("/post/{id}/delete", logging(deletePostTemplate))
	router.HandleFunc("/post/{id}/delete/post", logging(deletePostForm))
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
