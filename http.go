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
	EditID  int       `json:"EditID"`
	Content string    `json:"Content"`
	Title   string    `json:"Title"`
	Sucsess bool      `json:"Sucsess"`
	Time    time.Time `json:"Time"`
}

/*
	HTML:
*/

func displayIndex(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/index.html"))
	err := tmpl.ExecuteTemplate(w, "index", posts)
	if err != nil {
		log.Println(err)
	}
}
func requestPostWeb(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pubid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	result := readpostDB(pubid)
	post := postresp{PubID: pubid, Content: result.Content, Title: result.Title, Sucsess: result.Sucsess, Time: result.Time, EditID: result.EditID}
	//content := strings.Replace(post.Content, "\n", "<br>", -1)
	//post.Content = content
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	if post.Sucsess == false {
		tmpl.ExecuteTemplate(w, "notFound", post)
		return
	}
	tmpl.ExecuteTemplate(w, "display", post)
}
func createPostTemplateWeb(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/post.html"))
	err := tmpl.ExecuteTemplate(w, "createpost", nil)
	if err != nil {
		log.Println(err)
	}
}
func createPostWeb(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	newpost := postresp{Content: r.FormValue("Content"), Title: r.FormValue("Title")}
	rand.Seed(time.Now().UnixNano())
	newpost.PubID = genFromSeed()
	newpost.EditID = genFromSeed()
	createPostDB(newpost)
	url := fmt.Sprintf("/post/%v/request", newpost.PubID)
	http.Redirect(w, r, url, 302)
}
func handle404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/display.html", "front/layout.html")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.ExecuteTemplate(w, "404", nil)
	if err != nil {
		log.Println(err)
	}
}
func editpost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pubid, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	post := readpostDB(pubid)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))

	if editid == post.EditID {
		err := tmpl.ExecuteTemplate(w, "edit", post)
		if err != nil {
			log.Println(err)
		}
	} else {
		url := fmt.Sprintf("/post/%v/request", post.PubID)
		http.Redirect(w, r, url, 302)
	}

}
func edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pubid, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	post := readpostDB(pubid)
	if editid != post.EditID {
		url := fmt.Sprintf("/post/%v/request", post.PubID)
		http.Redirect(w, r, url, 302)
	}
	post.Content = r.FormValue("Content")
	post.Title = r.FormValue("Title")
	err := saveChanges(post)
	if err != nil {
		log.Println(err)
	}
	url := fmt.Sprintf("/post/%v/edit/%v", post.PubID, post.EditID)
	http.Redirect(w, r, url, 302)
}

/*
	JSON API:
*/

func postcounterAPI(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	json.NewEncoder(w).Encode(posts)
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
	newpost.EditID = genFromSeed()
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
