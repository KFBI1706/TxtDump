package main

import (
	"encoding/json"
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

type postdata struct {
	ID      int           `json:"ID"`
	EditID  int           `json:"EditID"`
	Content string        `json:"Content"`
	Md      template.HTML `json:"Md"`
	Title   string        `json:"Title"`
	Time    time.Time     `json:"Time"`
	Views   int           `json:"Views"`
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
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	html := parse(result.Content)
	result.Md = html
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
	newpost := postdata{Content: r.FormValue("Content"), Title: r.FormValue("Title")}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = genFromSeed()
	newpost.EditID = genFromSeed()
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
func editPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	post, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))

	if editid == post.EditID {
		err := tmpl.ExecuteTemplate(w, "edit", post)
		if err != nil {
			log.Println(err)
		}
	} else {
		url := fmt.Sprintf("/post/%v/request", post.ID)
		http.Redirect(w, r, url, 302)
	}

}
func edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	editid, err := strconv.Atoi(vars["editid"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	post, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	if editid != post.EditID {
		url := fmt.Sprintf("/post/%v/request", post.ID)
		http.Redirect(w, r, url, 302)
	}
	post.Content = r.FormValue("Content")
	post.Title = r.FormValue("Title")
	err = saveChanges(post)
	if err != nil {
		log.Println(err)
	}
	url := fmt.Sprintf("/post/%v/edit/%v", post.ID, post.EditID)
	http.Redirect(w, r, url, 302)
}
func deletePostWeb(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	exsistingpost, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	if exsistingpost.EditID != editid {
		return
	}
	err = deletepost(exsistingpost)
	if err != nil {
		log.Println(err)
	}
}
func documentation(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Println(err)
	}
	doc := parse(string(file))
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	err = tmpl.ExecuteTemplate(w, "doc", postdata{Md: doc})
	if err != nil {
		log.Println(err)
	}
}

/*
	JSON API:
*/

func editPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	editid, err := strconv.Atoi(vars["editid"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	exsistingpost, err := readpostDB(ID)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	newpost := postdata{}
	err = json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
	}
	if exsistingpost.EditID != editid {
		fmt.Fprintln(w, "Edit id did not match actual Edit id")
		return
	}
	newpost.ID = exsistingpost.ID
	newpost.EditID = exsistingpost.EditID
	err = saveChanges(newpost)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(newpost)
}
func deletePostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	editid, _ := strconv.Atoi(vars["editid"])
	exsistingpost, err := readpostDB(ID)
	if exsistingpost.EditID != editid {
		return
	}
	err = deletepost(exsistingpost)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "No such post")
	}
}
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
	result, err := readpostDB(i)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	err = incrementViewCounter(result.ID)
	if err != nil {
		log.Println(err)
	}
}
func createPostAPI(w http.ResponseWriter, r *http.Request) {
	newpost := postdata{}
	rand.Seed(time.Now().UnixNano())
	newpost.ID = genFromSeed()
	newpost.EditID = genFromSeed()
	err := json.NewDecoder(r.Body).Decode(&newpost)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "No data posted!")
		return
	}
	defer r.Body.Close()
	createPostDB(newpost)

	json.NewEncoder(w).Encode(newpost)
	r.Body.Close()
}
func routerWalk(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	t, err := route.GetPathTemplate()
	if err != nil {
		return err
	}
	fmt.Println(t)
	return nil
}
