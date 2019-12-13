package main

import (
	"fmt"
	"github.com/calbim/snippetbox/pkg/models"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("HIHIHIHIHIHI")
	snippet := models.Snippet{
		Title:   "Dummy Title",
		Content: "Dummy Content",
		Expires: time.Now().Add(1),
	}
	app.snippets.Insert(snippet.Title, snippet.Content, snippet.Expires.String())
	w.Write([]byte("Create a new snippet"))
}
