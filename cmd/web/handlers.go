package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	app := InitApp()
	w.Header().Add("Server", "GO")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.Logger.Error(
			err.Error(),
			"method",
			r.Method,
			"uri",
			r.URL.RequestURI(),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		app.Logger.Info("Success", "method", r.Method, "uri", r.URL.RequestURI())
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.Logger.Error(
			err.Error(),
			"method",
			r.Method,
			"uri",
			r.URL.RequestURI(),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create snippet here..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet here..."))
}
