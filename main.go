package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tsuru/config"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type User struct {
	Username string
	Emails   []string
	Friends  []string
}

var templatesPath = "views"
var templates map[string]*template.Template

func init() {
	config.ReadConfigFile("settings.yaml")
	//TODO: Error Checking
	basePath, _ := config.GetString("TEMPLATES:BASE")
	layoutsPath, _ := config.GetString("TEMPLATES:LAYOUTS")
	partialsPath, _ := config.GetString("TEMPLATES:PARTIALS")

	dir, _ := os.Getwd()
	templatesPath = filepath.Join(dir, basePath)
	fmt.Printf("Processing templates in %s\n", templatesPath)

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob(templatesPath + "/" + layoutsPath + "/*")
	if err != nil {
		log.Fatal(err)
	}

	partials, err := filepath.Glob(templatesPath + "/" + partialsPath + "/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		files := append(partials, layout)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	fmt.Printf("Rendering: %s\n", name)
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("Template %s does not exist.", name)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, name, data)
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "Jeff"}
	vars := make(map[string]interface{})
	vars["user"] = user
	renderTemplate(w, "index", vars)
}

func SeriesDisplayHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Gorilla display series " + string(id)))
}

func SeriesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display series list"))
}

func main() {
	r := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	r.HandleFunc("/", HomeHandler)

	s := r.Host("test.datistry.com").Subrouter()
	s.HandleFunc("/series", SeriesIndexHandler)
	s.HandleFunc("/series/{id:[0-9]+}", SeriesDisplayHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":3333", loggedRouter)
}
