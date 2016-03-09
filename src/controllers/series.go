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
		args := map[string]interface{}{"series": series}
		render(w, "series", "name", args)
	}
}

func SeriesIndexHandler(w http.ResponseWriter, r *http.Request) {
	args := make(map[string]interface{})

	if series, err := app.SeriesRepo.GetAllSeries(); err == nil {
		args["series"] = series
	}

	render(w, "series", "index", args)
}

func SeriesIdHandler(w http.ResponseWriter, r *http.Request) {
	//BUG(archeious): If the series name ACTUALLY has a dash this will fail
	var series series.Series
	var err error
	vars := mux.Vars(r)
	if series, err = app.SeriesRepo.GetSeriesById(vars["id"]); err != nil {
		if err == item.ErrDoesNotExist {
			render(w, "error", "doesnotexist", nil)
		} else {
			w.Write([]byte("500 " + err.Error()))
		}
	} else {
		args := map[string]interface{}{"series": series}
		render(w, "series", "name", args)
	}
}

func SeriesEditFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if series, err := app.SeriesRepo.GetSeriesById(vars["urlid"]); err != nil {
		w.Write([]byte("500 " + err.Error()))
		return
	} else {
		args := map[string]interface{}{"series": series}
		render(w, "series", "edit", args)
	}
}

func SeriesEditHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Error Processing Form")
		w.Write([]byte("500 " + err.Error()))
		return
	}
	vars := mux.Vars(r)
	urlid := vars["urlid"]
	id := r.FormValue("id")
	if urlid != id {
		fmt.Println("ID MISMATCH:", urlid, id)
		return
	}
	series, err := app.SeriesRepo.GetSeriesById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	series.SetName(r.FormValue("name"))
	series.SetJName(r.FormValue("jname"))
	series.SetSynopsis(r.FormValue("synopsis"))
	series.SetExtID("imdb", r.FormValue("imdbid"))
	series.SetExtID("mal", r.FormValue("malid"))
	app.SeriesRepo.SaveSeries(series)
	args := map[string]interface{}{"series": series}
	render(w, "series", "edit", args)
}

func SeriesAddFormHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "series", "add", nil)
}

func SeriesAddHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error processing form")
	}

	name := r.FormValue("name")

	//TODO: Error Check
	//TODO: nil species
	if newSeries, err := app.SeriesRepo.NewSeries(name); err == nil {
		args := make(map[string]interface{})
		if extid := r.FormValue("malid"); extid != "" {
			newSeries.SetExtID("mal", extid)
		}
		if extid := r.FormValue("imdbid"); extid != "" {
			newSeries.SetExtID("imdb", extid)
		}
		if synopsis := r.FormValue("synopsis"); synopsis != "" {
			newSeries.SetSynopsis(synopsis)
		}

		if err := app.SeriesRepo.SaveSeries(newSeries); err != nil {
			fmt.Println(err)
		}
		args["imdbId"] = r.FormValue("imdbid")
		args["action"] = "add"
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
