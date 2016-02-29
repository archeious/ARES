package controllers

import (
	"app"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"models/item"
	"models/series"
	"net/http"
	"strings"
)

//TODO: BUG: The the series name ACTUALLY has a dash this will fail
func SeriesNameHandler(w http.ResponseWriter, r *http.Request) {
	var series series.Series
	var err error
	vars := mux.Vars(r)
	name := strings.Replace(vars["name"], "-", " ", -1)
	if series, err = app.SeriesRepo.GetSeriesByName(name); err != nil {
		if err == item.ErrDoesNotExist {
			render(w, "error", "doesnotexist", nil)
		} else {
			w.Write([]byte("500 " + err.Error()))
		}
	} else {
		args := map[string]interface{}{"name": series.Name(), "user": "Jeff", "series": series}
		render(w, "series", "name", args)
	}
}

func SeriesIndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout.tmpl", "templates/series/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	args := make(map[string]string)
	args["title"] = "Test Title"

	if err := t.Execute(w, args); err != nil {
		log.Fatal(err)
	}
}
