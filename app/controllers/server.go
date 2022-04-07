package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/hiromi-mitsuoka/golang-todo/app/models"
	"github.com/hiromi-mitsuoka/golang-todo/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// NOTE: Cache templates in advance
	templates := template.Must(template.ParseFiles(files...))
	// NOTE: Read the file declared by {{define "layout"}}
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// NOTE: A cookie can be retrieved from the request
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		// Check if there is a record in the DB
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			// NOTE: 404 page not found
			http.NotFound(w, r)
			return
		}
		// strconv.Atoi: String convert string to int
		qi, err := strconv.Atoi(q[2]) // Retrieve todo's id
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	// NOTE: Use if you are loading css, js
	// files := http.FileServer(http.Dir(config.Config.Static))
	// http.Handle("/static/", http.StripPrefix("/static", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit)) // If it ends in /, consider that the id is included
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	// ListenAndServe: (https://zenn.dev/hsaki/books/golang-httpserver-internal/viewer/serverstart)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
