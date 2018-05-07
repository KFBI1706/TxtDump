package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/KFBI1706/Txtdump/api"
	"github.com/KFBI1706/Txtdump/html"

	"github.com/KFBI1706/Txtdump/helper"
	"github.com/KFBI1706/Txtdump/sql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func main() {
	err := sql.TestDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	test := model.PostMeta{PostPerms: 2}
	dbdrop := flag.Bool("dropdb", false, "Drop current table and all data")
	dbsetup := flag.Bool("setupdb", false, "Setup db when running")
	production := flag.Bool("production", false, "Is the server in production?")
	port := flag.Int("port", 1337, "for using a custom port")
	flag.Parse()
	addr := fmt.Sprintf(":%v", *port)
	if *dbdrop || *dbsetup {
		if *dbdrop {
			err = helper.ClearOutDB()
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println(*dbsetup)
		if *dbsetup {
			err = helper.SetupDB()
			if err != nil {
				log.Println(err)
			}
		}
		os.Exit(3)
	}
	log.Printf("%v Post(s) Currently in DB\n", sql.CountPosts())
	CSRF := csrf.Protect([]byte("OTAyNDhmajBkYnBhamtudnBhc29ldXI"), csrf.Secure(*production))
	router := mux.NewRouter()
	router.HandleFunc("/", logging(html.DisplayIndex)).Methods("GET")
	router.HandleFunc("/api/v1/post/amount", logging(api.PostcounterAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(api.RequestPostAPI)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}/request", logging(api.RequestPostWithPassAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/create", logging(api.CreatePostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/edit", logging(api.EditPostAPI)).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/delete", logging(api.DeletePostAPI))
	router.HandleFunc("/post/{id}/request", logging(html.RequestPostWeb)).Methods("GET")
	router.HandleFunc("/post/{id}/request/decrypt", logging(html.RequestPostDecrypt)).Methods("POST")
	router.HandleFunc("/post/{id}/edit", logging(html.EditPostTemplate))
	router.HandleFunc("/post/{id}/edit/decrypt", logging(html.EditPostDecrypt))
	router.HandleFunc("/post/{id}/edit/post", logging(html.EditPostForm)).Methods("POST")
	router.HandleFunc("/post/{id}/delete", logging(html.DeletePostTemplate))
	router.HandleFunc("/post/{id}/delete/post", logging(html.DeletePostForm))
	router.HandleFunc("/post/create", logging(html.CreatePostTemplateWeb))
	router.HandleFunc("/post/create/new", logging(html.CreatePostWeb))
	router.HandleFunc("/documentation", logging(html.Documentation))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	router.Walk(routerWalk)
	err = http.ListenAndServe(addr, CSRF(router))
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
