package controllers

import (
	"app"
	"models/item"
	"net/http"
)

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	if series, err := app.SeriesRepo.GetSeriesByTag("featured"); err != nil {
		if err == item.ErrDoesNotExist {
			render(w, "error", "doesnotexist", nil)
		} else {
			w.Write([]byte("500 " + err.Error()))
		}
	} else {
		args := map[string]interface{}{"series": series}
		render(w, "site", "frontpage", args)
	}
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "site", "about", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "site", "contact", nil)
}
