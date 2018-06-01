package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KFBI1706/TxtDump/api"
	"github.com/KFBI1706/TxtDump/config"
	"github.com/KFBI1706/TxtDump/html"
	"github.com/KFBI1706/TxtDump/sql"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.ParseConfig("development")
	var defaultDir string
	if conf.DBStringLocation != "" {
		defaultDir = conf.DBStringLocation
	} else {
		defaultDir = "dbstring"
	}
	dir := flag.String("dir", defaultDir, "root-directory for important files such as dbstring")
	dbdrop := flag.Bool("dropdb", false, "Drop current table and all data")
	dbsetup := flag.Bool("setupdb", false, "Setup db when running")
	//production flag should be an explicit env variable
	port := flag.Int("port", conf.Port, "for using a custom port")
	flag.Parse()
	config.InitDB(*dir)
	defer config.DB.Close()
	addr := fmt.Sprintf(":%v", *port)
	if *dbdrop || *dbsetup {
		if *dbdrop {
			err := sql.ClearOutDB()
			if err != nil {
				log.Println(err)
			}
		}
		if *dbsetup {
			err := sql.SetupDB()
			if err != nil {
				log.Println(err)
			}
		}
		os.Exit(3)
	}
	err := sql.TestDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v Post(s) Currently in DB\n", sql.CountPosts())
	router := mux.NewRouter()
	a := router.PathPrefix("/api/v1/post").Subrouter()
	w := router.PathPrefix("/post/").Subrouter()
	router.HandleFunc("/", logging(html.DisplayIndex)).Methods("GET")
	a.HandleFunc("/amount", logging(api.PostcounterAPI)).Methods("GET")
	a.HandleFunc("/{id}/request", logging(api.RequestPostAPI)).Methods("GET")
	a.HandleFunc("/{id}/request", logging(api.RequestPostWithPassAPI)).Methods("POST")
	a.HandleFunc("/create", logging(api.CreatePostAPI)).Methods("POST")
	a.HandleFunc("/{id}/edit", logging(api.EditPostAPI)).Methods("POST")
	a.HandleFunc("/{id}/delete", logging(api.DeletePostAPI)).Methods("POST")
	w.HandleFunc("/{id}/request", logging(html.RequestPostWeb)).Methods("GET")
	w.HandleFunc("/{id}/request/decrypt", logging(html.RequestPostDecrypt)).Methods("POST")
	w.HandleFunc("/{id}/edit", logging(html.EditPostTemplate))
	w.HandleFunc("/{id}/edit/decrypt", logging(html.EditPostDecrypt))
	w.HandleFunc("/{id}/edit/post", logging(html.EditPostForm)).Methods("POST")
	w.HandleFunc("/{id}/delete", logging(html.DeletePostTemplate))
	w.HandleFunc("/{id}/delete/post", logging(html.DeletePostForm)).Methods("POST")
	w.HandleFunc("/create", logging(html.CreatePostTemplateWeb))
	w.HandleFunc("/create/new", logging(html.CreatePostWeb))
	router.HandleFunc("/documentation", logging(html.Documentation))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("front/"))))
	err = router.Walk(routerWalk)
	if err != nil {
		log.Fatal(err)
	}
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
