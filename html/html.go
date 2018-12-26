package html

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/KFBI1706/TxtDump/crypto"
	"github.com/KFBI1706/TxtDump/helper"
	"github.com/KFBI1706/TxtDump/model"
	"github.com/KFBI1706/TxtDump/sql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

//ProcessRequest Processes an request for post
func ProcessRequest(w http.ResponseWriter, r *http.Request) model.Post {
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

//DisplayIndex renders the Index template with some metadata
func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/index.html"))
	metas, err := sql.PostMetas()
	if err != nil {
		log.Println(err)
	}
	datas, err := sql.PostDatas()
	if err != nil {
		log.Println(err)
	}
	err = tmpl.ExecuteTemplate(w, "index", model.M{"Count": sql.CountPosts(), "Meta": metas, "Data": datas})
	if err != nil {
		log.Println(err)
	}
}

func parsePost(ID int, data model.Data) (markdown model.Markdown) {
	defer sql.IncrementViewCounter(ID)
	markdown.MD = parse(data.Content)
	if mdhead := getMDHeader(markdown.MD); mdhead != "" && data.Title == "" {
		data.Title = mdhead
	}
	if thumb := getIMG(markdown.MD); thumb != "" {
		markdown.IMG = template.HTML(thumb)
	}
	markdown.TitleMD = template.HTML(data.Title)
	return
}

//RequestPostDecrypt does exactly like what it sounds like it does
func RequestPostDecrypt(w http.ResponseWriter, r *http.Request) {
	post := ProcessRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	post.Crypto.Hash = r.FormValue("Pass")
	if crypto.RequestDecrypt(&post) {
		err := tmpl.ExecuteTemplate(w, "display", model.M{"ID": post.ID, "Markdown": parsePost(post.ID, post.Data),
			"Meta": post.Meta})
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "Something went wrong")
			return
		}
	}
}

//RequestPostWeb renders the content of the post to W/http.ResponseWriter
func RequestPostWeb(w http.ResponseWriter, r *http.Request) {
	post := ProcessRequest(w, r)
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	if post.Data.PostPerms <= 2 {
		err := tmpl.ExecuteTemplate(w, "display", model.M{"ID": post.ID, "markdown": parsePost(post.ID, post.Data),
			"meta": post.Meta})
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "Something went wrong")
			return
		}
	} else {
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
	rand.Seed(time.Now().UnixNano())
	newpost := model.Post{ID: helper.GenFromSeed(), Data: model.Data{Content: r.FormValue("Content"), Title: r.FormValue("Title")}}
	newpost.Data.PostID = newpost.ID
	newpost.Meta.PostID = newpost.ID
	newpost.Data.PostPerms, err = helper.DeterminePerms(r.FormValue("postperms"))
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
		if post.Data.PostPerms == 3 {
			err := tmpl.ExecuteTemplate(w, "displayPass", map[string]interface{}{
				"ID":   post.ID,
				"Mode": templateString,
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			keys, ok := r.URL.Query()["failed"]
			if ok && len(keys) > 0 && (len(r.URL.Query()["title"]) > 0 && len(r.URL.Query()["content"]) > 0) {
				post.Data.Title = r.URL.Query()["title"][0]
				post.Data.Content = r.URL.Query()["content"][0]
			}
			err := tmpl.ExecuteTemplate(w, "edit", map[string]interface{}{
				"ID":      post.ID,
				"Title":   post.Data.Title,
				"Content": post.Data.Content,
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
	post.Crypto.Hash = r.FormValue("Pass")
	if crypto.RequestDecrypt(&post) {
		parsePost(post.ID, post.Data)
		err := tmpl.ExecuteTemplate(w, "edit", map[string]interface{}{
			csrf.TemplateTag: csrf.TemplateField(r),
			"ID":             post.ID,
			"Title":          post.Data.Title,
			"Content":        post.Data.Content,
		})
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "Something went wrong")
			return
		}
	}
}

//EditPostTemplate is a handler function used to call postTemplate with the "edit" parameter
func EditPostTemplate(w http.ResponseWriter, r *http.Request) {
	postTemplate(w, r, "edit")
}

//DeletePostTemplate is a handler function used to call postTemplate with the "delete" parameter
func DeletePostTemplate(w http.ResponseWriter, r *http.Request) {
	postTemplate(w, r, "deletepost")
}

func postForm(w http.ResponseWriter, r *http.Request, operation string) {
	post := ProcessRequest(w, r)
	//TODO: rewrite this, not broken, just bad
	url := fmt.Sprintf("/post/%v/request", post.ID)
	var err error
	if operation == "edit" {
		post.Data.Content = r.FormValue("Content")
		post.Data.Title = r.FormValue("Title")
		hash := post.Crypto.Hash
		post.Crypto.Hash = r.FormValue("Pass")
		if post.Data.PostPerms == 3 {
			key := crypto.GetEncKey(&post)
			fmt.Println("encrypting", post.Data.Content)
			b := []byte(post.Data.Content)
			ct, err := crypto.EncryptBytes(b, &key)
			if err != nil {
				log.Fatal(err)
			}
			post.Data.Content = b64.StdEncoding.EncodeToString(ct)
		}
		ok := crypto.CheckPass(post.Crypto.Hash, post.ID, post.Data.PostPerms)
		log.Printf("Pass check %v\n Pass values; hash: %s id: %d perms: %d\n", ok, post.Crypto.Hash, post.ID, post.Data.PostPerms)
		if ok {
			post.Crypto.Hash = hash
			err = sql.SaveChanges(post)
			if err != nil {
				log.Println(err)
			}
		} else {
			url = fmt.Sprintf("/post/%v/edit?failed=true&title=%v&content=%v", post.ID, post.Data.Title, post.Data.Content)
		}
	} else if operation == "delete" {
		post.Crypto.Hash = r.FormValue("Pass")
		if crypto.CheckPass(post.Crypto.Hash, post.ID, post.Data.PostPerms) {
			err = sql.DeletePost(post)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", 302)
			return
		}
	}
	if err != nil {
		fmt.Fprintln(w, "Something went wrong")
		return
	}
	http.Redirect(w, r, url, 302)
}

//EditPostForm handle function to call postForm with the "edit" parameter
func EditPostForm(w http.ResponseWriter, r *http.Request) {
	postForm(w, r, "edit")
}

//DeletePostForm handle function to call postForm with the "delete" parameter
func DeletePostForm(w http.ResponseWriter, r *http.Request) {
	postForm(w, r, "delete")
}

//Documentation used to display the README file as documentation
func Documentation(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("front/layout.html", "front/display.html"))
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.ExecuteTemplate(w, "doc", model.Markdown{MD: parse(string(file))})
	if err != nil {
		log.Println(err)
	}
}
