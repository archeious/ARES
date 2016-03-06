package controllers

import (
	"app"
	"fmt"
	"github.com/gorilla/mux"
	//	"html/template"
	//	"log"
	"models/item"
	"models/series"
	"net/http"
	"strings"
)

func SeriesNameHandler(w http.ResponseWriter, r *http.Request) {
	//BUG(archeious): If the series name ACTUALLY has a dash this will fail
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
	args := make(map[string]interface{})
	args["user"] = "Jeff"
	args["title"] = "Test Title"
	fmt.Println(args)

	if series, err := app.SeriesRepo.GetAllSeries(); err == nil {
		args["series"] = series
	}

	fmt.Println(args)
	render(w, "series", "index", args)
}

func SeriesIdHandler(w http.ResponseWriter, r *http.Request) {
	//BUG(archeious): If the series name ACTUALLY has a dash this will fail
	var series series.Series
	var err error
	vars := mux.Vars(r)
	fmt.Println(vars)
	if series, err = app.SeriesRepo.GetSeriesById(vars["id"]); err != nil {
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

func SeriesAddFormHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "series", "add", nil)
}

func SeriesAddHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error processing form")
	}
	fmt.Println(r.Form)

	name := r.FormValue("name")
	//TODO: Error Check
	//TODO: nil species
	newSeries, _ := app.SeriesRepo.NewSeries(name, "")

	args := make(map[string]interface{})
	args["series"] = newSeries
	vars := mux.Vars(r)
	fmt.Println(vars)
	fmt.Println(args)
	render(w, "series", "index", nil)
}
