package controllers

import (
	"app"
	"models/item"
	"net/http"

	"github.com/gorilla/context"
)

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user")
	if series, err := app.SeriesRepo.GetSeriesByTag("featured"); err != nil {
		if err == item.ErrDoesNotExist {
			render(w, r, "error", "doesnotexist", nil)
		} else {
			w.Write([]byte("500 " + err.Error()))
		}
	} else {
		args := map[string]interface{}{"series": series, "user": u}
		render(w, r, "site", "frontpage", args)
	}
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "site", "about", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "site", "contact", nil)
}
