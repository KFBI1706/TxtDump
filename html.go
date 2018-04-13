package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type postData struct {
	ID      int           `json:"ID"`
	EditID  string        `json:"EditID"`
	Content string        `json:"Content"`
	Md      template.HTML `json:""`
	Title   string        `json:"Title"`
	TitleMD template.HTML `json:""`
	Time    time.Time     `json:"Time"`
	Views   int           `json:"Views"`
}

func displayIndex(w http.ResponseWriter, r *http.Request) {
	posts := postcounter{Count: countPosts()}
	posts, err := postMeta()
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/index.html"))
	err = tmpl.ExecuteTemplate(w, "index", posts)
	if err != nil {
		log.Println(err)
	}
}
func requestPostWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Request needs to be int")
		return
	}
	result, err := readpostDB(ID)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	if err != nil {
		log.Println(err)
		tmpl.ExecuteTemplate(w, "notFound", result)
		return
	}
	result.Md = parse(result.Content)
	mdhead := getMDHeader(result.Md)
	if mdhead != "" && result.Title == "" {
		result.Title = mdhead
	}
	result.TitleMD = template.HTML(result.Title)
	tmpl.ExecuteTemplate(w, "display", result)
	err = incrementViewCounter(result.ID)
	if err != nil {
		log.Println(err)
	}
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
	newpost := postData{Content: r.FormValue("Content"), Title: r.FormValue("Title")}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = genFromSeed()
	newpost.EditID = string(genFromSeed())
	if r.FormValue("CustomPass") == "on" {
		sec := r.FormValue("Pass")
		newpost.EditID = sec
	}
	createPostDB(newpost)
	url := fmt.Sprintf("/post/%v/request", newpost.ID)
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
func editPostTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	post, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/post.html"))
	err = tmpl.ExecuteTemplate(w, "edit", post)
	if err != nil {
		log.Println(err)
	}

}
func editPostForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	post, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	post.Content = r.FormValue("Content")
	post.Title = r.FormValue("Title")
	post.EditID = r.FormValue("Pass")
	err = saveChanges(post)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Wrong Password")
		return
	}
	url := fmt.Sprintf("/post/%v/request", post.ID)
	http.Redirect(w, r, url, 302)
}
func deletePostTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postid, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/post.html"))
	post := postData{ID: postid}
	err = tmpl.ExecuteTemplate(w, "deletepost", post)
	if err != nil {
		log.Println(err)
	}
}
func deletePostForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	post, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	post.EditID = r.FormValue("Pass")
	err = deletepost(post)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Wrong Password")
		return
	}
}
func documentation(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Println(err)
	}
	doc := parse(string(file))
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	err = tmpl.ExecuteTemplate(w, "doc", postData{Md: doc})
	if err != nil {
		log.Println(err)
	}
}
