package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//The serveError helper writes an error message and stack trace to the errorLog,
//then sends a generic 500 Internal Server Error response to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Output(2, fmt.Sprintf("%s \n %s", err.Error(), debug.Stack()))
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

//The clientError helper sends a specific error code to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
