package html

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/KFBI1706/TxtDump/crypto"
	"github.com/KFBI1706/TxtDump/helper"
	"github.com/KFBI1706/TxtDump/model"
	"github.com/KFBI1706/TxtDump/sql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

//DisplayIndex renders the Index template with some metadata
func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	posts := model.PostCounter{Count: sql.CountPosts()}
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/index.html"))

	posts, err := sql.PostMetas()
	if err != nil {
		log.Println(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", posts)
	if err != nil {
		log.Println(err)
	}
}

//ProcessRequest Processes an request for post
func ProcessRequest(w http.ResponseWriter, r *http.Request) model.PostData {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Request needs to be int")
	}
	post, err := sql.ReadPostDB(id)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "ID not found")
	}
	return post
}

func parsePost(post *model.PostData) {
	defer sql.IncrementViewCounter(post.ID)
	post.Md = parse(post.Content)
	mdhead := getMDHeader(post.Md)
	if mdhead != "" && post.Title == "" {
		post.Title = mdhead
	}
	post.TitleMD = template.HTML(post.Title)
}

//RequestPostDecrypt does exactly like what it sounds like it does
func RequestPostDecrypt(w http.ResponseWriter, r *http.Request) {
	post := ProcessRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	post.Hash = r.FormValue("Pass")
	if crypto.RequestDecrypt(&post) {
		parsePost(&post)
		tmpl.ExecuteTemplate(w, "display", post)
	}
}

//RequestPostWeb renders the content of the post to W/http.ResponseWriter
func RequestPostWeb(w http.ResponseWriter, r *http.Request) {
	post := ProcessRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	if post.PostPerms == 1 || post.PostPerms == 2 {
		parsePost(&post)
		tmpl.ExecuteTemplate(w, "display", post)
	} else if post.PostPerms == 3 {
		err := tmpl.ExecuteTemplate(w, "displayPass", map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
			"ID":             post.ID,
			"Mode":           "request",
		})
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "Something went wrong")
			return
		}
	}
}

//CreatePostTemplateWeb renders the html template for adding a new post to W/http.ResponseWriter
func CreatePostTemplateWeb(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/post.html"))
	err := tmpl.ExecuteTemplate(w, "createpost", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
	if err != nil {
		log.Println(err)
	}
}

//CreatePostWeb pareses posted data from r and registers it to DB
func CreatePostWeb(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	newpost := model.PostData{Content: r.FormValue("Content"), Title: r.FormValue("Title")}
	newpost.PostPerms, err = helper.DeterminePerms(r.FormValue("postperms"))
	if err != nil {
		log.Println(err)
	}
	crypto.SecurePost(&newpost, r.FormValue("Pass"))
	sql.CreatePostDB(newpost)
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

func postTemplate(w http.ResponseWriter, r *http.Request, templateString string) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html", "front/post.html"))
	post := ProcessRequest(w, r)
	if templateString == "edit" {
		if post.PostPerms == 3 {
			err := tmpl.ExecuteTemplate(w, "displayPass", map[string]interface{}{
				csrf.TemplateTag: csrf.TemplateField(r),
				"ID":             post.ID,
				"Mode":           templateString,
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			err := tmpl.ExecuteTemplate(w, "edit", map[string]interface{}{
				csrf.TemplateTag: csrf.TemplateField(r),
				"ID":             post.ID,
				"Title":          post.Title,
				"Content":        post.Content,
			})
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		err := tmpl.ExecuteTemplate(w, templateString, map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
			"ID":             post.ID,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

//EditPostDecrypt renders the edit template
func EditPostDecrypt(w http.ResponseWriter, r *http.Request) {
	post := ProcessRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html", "front/post.html"))

	//post.Hash = r.FormValue("Pass")
	//if crypto.RequestDecrypt(&post) {
	//	parsePost(&post)
	//	tmpl.ExecuteTemplate(w, "display", post)
	//}
	post.Hash = r.FormValue("Pass")
	if crypto.RequestDecrypt(&post) {
		parsePost(&post)
		log.Println("parsed post")
		log.Println(post.ID)
		log.Println(post.Title)
		log.Println(post.Content)
		err := tmpl.ExecuteTemplate(w, "edit", map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
			"ID":             post.ID,
			"Title":          post.Title,
			"Content":        post.Content,
		})
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "Something went wrong")
			return
		}
	}
}

func EditPostTemplate(w http.ResponseWriter, r *http.Request) {
	postTemplate(w, r, "edit")
}

func DeletePostTemplate(w http.ResponseWriter, r *http.Request) {
	postTemplate(w, r, "deletepost")
}

func postForm(w http.ResponseWriter, r *http.Request, operation string) {
	post := ProcessRequest(w, r)
	var err error
	if operation == "edit" {
		post.Content = r.FormValue("Content")
		post.Title = r.FormValue("Title")
		post.Hash = r.FormValue("Pass")
		if post.PostPerms == 3 {
			key := crypto.GetEncKey(&post)
			b := []byte(post.Content)
			ct, err := crypto.EncryptBytes(b, &key)
			if err != nil {
				log.Fatal(err)
			}
			post.Content = b64.StdEncoding.EncodeToString(ct)
		}
		if valid := crypto.CheckPass(post.Hash, post.ID, post.PostPerms); valid {
			err = sql.SaveChanges(post)
		}
	} else if operation == "delete" {
		post.Hash = r.FormValue("Pass")
		err = sql.DeletePost(post)
	}
	url := "/"
	if err != nil {
		log.Println(err)
		url = fmt.Sprintf("/post/%v/request", post.ID)
	}
	http.Redirect(w, r, url, 302)

}

//EditPostForm handles
func EditPostForm(w http.ResponseWriter, r *http.Request) {
	postForm(w, r, "edit")
}

func DeletePostForm(w http.ResponseWriter, r *http.Request) {
	postForm(w, r, "delete")
}
func Documentation(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Println(err)
	}
	doc := parse(string(file))
	err = tmpl.ExecuteTemplate(w, "doc", model.PostData{Md: doc})
	if err != nil {
		log.Println(err)
	}
}
