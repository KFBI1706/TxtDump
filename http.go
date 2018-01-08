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
	ID      int    `json:"-"`
	PubID   int    `json:"PubID"`
	Content string `json:"Content"`
	Title   string `json:"Title"`
	Sucsess bool   `json:"Sucsess"`
	Time    string `json:"Time"`
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
func createPost(w http.ResponseWriter, r *http.Request) {
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
