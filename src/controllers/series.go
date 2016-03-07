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
	fmt.Println("FORM:", r.Form)

	name := r.FormValue("name")

	//TODO: Error Check
	//TODO: nil species
	if newSeries, err := app.SeriesRepo.NewSeries(name, ""); err == nil {
		args := make(map[string]interface{})
		if extid := r.FormValue("malid"); extid != "" {
			newSeries.SetExtID("mal", extid)
		}
		if extid := r.FormValue("imdbid"); extid != "" {
			newSeries.SetExtID("imdb", extid)
		}
		if err := app.SeriesRepo.SaveSeries(newSeries); err != nil {
			fmt.Println(err)
		}
		args["imdbId"] = r.FormValue("imdbid")

		fmt.Println("ARGS:", args)
		render(w, "series", "index", nil)
	} else {
		if err == item.ErrAlreadyExists {
			http.Redirect(w, r, "/series/edit/"+newSeries.Id(), http.StatusFound)
			return
		} else {
			fmt.Println(err)
		}
	}
}
