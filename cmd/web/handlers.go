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
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
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
	s, err := app.snippets.Get(id)
	if err == models.ErrRecordNotFound {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	app.snippets.Insert("Dummy Title", "Dummy Content", time.Now().Add(1).String())

	http.Redirect(w, r, "/snippet?id=%d", http.StatusSeeOther)
}
