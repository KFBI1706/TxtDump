package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type postData struct {
	ID        int           `json:"ID"`
	EditID    int           `json:"EditID"`
	Hash      string        `json:"Hash"`
	Salt      string        `json:"Salt"`
	authHash  string        `json:authHash`
	Key       string        `json:Key`
	PostPerms int           `json:"PostPerms,string"`
	Content   string        `json:"Content"`
	Md        template.HTML `json:""`
	Title     string        `json:"Title"`
	TitleMD   template.HTML `json:""`
	Time      time.Time     `json:"Time"`
	Views     int           `json:"Views"`
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

func processRequest(w http.ResponseWriter, r *http.Request) (post postData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	post, err = readpostDB(id)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Something went wrong")
	}
	return
}

func requestPostDecrypt(w http.ResponseWriter, r *http.Request) {
	post := processRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	post.Hash = r.FormValue("Pass")
	fmt.Println("request for decryption: " + post.Hash)
	if requestDecrypt(&post) {
		err := incrementViewCounter(post.ID)
		if err != nil {
			log.Println(err)
		}
		post.Md = parse(post.Content)
		mdhead := getMDHeader(post.Md)
		if mdhead != "" && post.Title == "" {
			post.Title = mdhead
		}
		post.TitleMD = template.HTML(post.Title)
		if post.PostPerms == 1 || post.PostPerms == 2 {
			tmpl.ExecuteTemplate(w, "display", post)
			err = incrementViewCounter(post.ID)
			if err != nil {
				log.Println(err)
			}
		}
		tmpl.ExecuteTemplate(w, "display", post)
	} else {
		http.Redirect(w, r, "/", 302)
	}

}

func requestPostWeb(w http.ResponseWriter, r *http.Request) {
	result := processRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	result.Md = parse(result.Content)
	mdhead := getMDHeader(result.Md)
	if mdhead != "" && result.Title == "" {
		result.Title = mdhead
	}
	result.TitleMD = template.HTML(result.Title)
	if result.PostPerms == 1 || result.PostPerms == 2 {
		tmpl.ExecuteTemplate(w, "display", result)
		err := incrementViewCounter(result.ID)
		if err != nil {
			log.Println(err)
		}
	} else if result.PostPerms == 3 {
		tmpl.ExecuteTemplate(w, "displayPass", result)
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
	newpost.PostPerms, err = determinePerms(r.FormValue("postperms"))
	if err != nil {
		log.Println(err)
	}
	securePost(&newpost, r.FormValue("Pass"))
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
	postTemplate(w, r, "edit")
}

func deletePostTemplate(w http.ResponseWriter, r *http.Request) {
	postTemplate(w, r, "deletepost")
}

func postTemplate(w http.ResponseWriter, r *http.Request, templateString string) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/post.html"))
	post := processRequest(w, r)
	err := tmpl.ExecuteTemplate(w, templateString, post)
	if err != nil {
		log.Println(err)
	}
}

func editPostForm(w http.ResponseWriter, r *http.Request) {
	post := processRequest(w, r)
	post.Content = r.FormValue("Content")
	post.Title = r.FormValue("Title")
	post.Hash = r.FormValue("Pass")
	err := saveChanges(post)
	url := "/"
	if err != nil {
		log.Println(err)
	} else {
		url = fmt.Sprintf("/post/%v/request", post.ID)
	}
	http.Redirect(w, r, url, 302)
}
func deletePostForm(w http.ResponseWriter, r *http.Request) {
	post := processRequest(w, r)
	post.Hash = r.FormValue("Pass")
	err := deletepost(post)
	url := "/"
	if err != nil {
		log.Println(err)
	} else {
		url = fmt.Sprintf("/post/%v/request", post.ID)
	}
	http.Redirect(w, r, url, 302)
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
