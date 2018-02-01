package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type postresp struct {
	ID      int       `json:"-"`
	PubID   int       `json:"PubID"`
	Content string    `json:"Content"`
	Title   string    `json:"Title"`
	Sucsess bool      `json:"Sucsess"`
	Time    time.Time `json:"Time"`
}
type postcounter struct {
	Count   int      `json:"Count"`
	PostIDs []int    `json:"ID"`
	Titles  []string `json:"Titles"`
}

func displayIndex(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/index.html"))
	err := tmpl.ExecuteTemplate(w, "index", posts)
	if err != nil {
		log.Println(err)
	}
}
func postcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	json.NewEncoder(w).Encode(posts)
}
func requestPostHTML(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	result := readpostDB(i)
	post := postresp{PubID: i, Content: result.Content, Title: result.Title, Sucsess: result.Sucsess, Time: result.Time}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	if post.Sucsess == false {
		tmpl.ExecuteTemplate(w, "notFound", post)
		return
	}
	tmpl.ExecuteTemplate(w, "display", post)
}
func requestPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	result := readpostDB(i)
	post := postresp{PubID: i, Content: result.Content, Title: result.Title, Sucsess: true, Time: result.Time}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(post)
}
func createPostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := postresp{}
	rand.Seed(time.Now().UnixNano())
	newpost.PubID = genFromSeed()
	log.Print(newpost.PubID)
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		fmt.Fprint(w, "No data posted!")
		return
	}
	defer r.Body.Close()
	createPostDB(newpost)
	newpost.Sucsess = true
	json.NewEncoder(w).Encode(newpost)
}
func createPostWeb(w http.ResponseWriter, r *http.Request) {
	newpost := postresp{}
	rand.Seed(time.Now().UnixNano())
	newpost.PubID = genFromSeed()
	log.Print(newpost.PubID)
}
func handle404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/display.html", "front/layout.html")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.Execute(w, "404")
	if err != nil {
		log.Println(err)
	}
}
