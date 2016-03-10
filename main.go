package main

import (
	"controllers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"github.com/tsuru/config"
	"html/template"
	"log"
	"models/user"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	templatesPath = "views"
	templates     map[string]*template.Template
	UserRepo      user.UserRepository
)

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

	UserRepo = user.NewBaseUserRepository()
	if _, err := UserRepo.NewUser("jeff", "password"); err != nil {
		fmt.Println(err)
	}
}

func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("Template %s does not exist.", name)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user := user.BaseUser{Username: "Jeff"}
	vars := make(map[string]interface{})
	vars["user"] = user
	renderTemplate(w, "index", vars)
}

func DisplayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Login.html", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := controllers.Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if _, ok := session.Values["user"]; ok {
		delete(session.Values, "user")
		if err := session.Save(r, w); err != nil {
			fmt.Println("Save Error:", err)
		}
	}
	http.Redirect(w, r, "/", 302)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := strings.Join(r.Form["username"], "")
	password := strings.Join(r.Form["password"], "")

	fmt.Println("username:", username, "password:", password)

	u, _ := UserRepo.GetUserByName(username)

	fmt.Println("user obj:", u)

	if ok, err := u.ValidatePassword(password); ok {
		fmt.Println("password validated")
		session, err := controllers.Store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		session.Values["user"] = username
		if err := session.Save(r, w); err != nil {
			fmt.Println("Save Error:", err)
		}
		http.Redirect(w, r, "/", 302)
	} else {
		fmt.Println(err)
		renderTemplate(w, "Login.html", nil)
	}
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
	userAuth := controllers.UserAuthMiddleWare(loggedRouter)
	recovery := handlers.RecoveryHandler()(userAuth)

	s := r.Host("test.datistry.com").Subrouter()
	s.HandleFunc("/", controllers.FrontPageHandler)
	s.HandleFunc("/about", controllers.AboutPageHandler)
	s.HandleFunc("/contact", controllers.ContactPageHandler)
	s.HandleFunc("/login", DisplayLoginHandler).Methods("GET")
	s.HandleFunc("/login", LoginHandler).Methods("POST")
	s.HandleFunc("/logout", LogoutHandler)
	s.HandleFunc("/series/edit/{urlid:[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+}", controllers.SeriesEditFormHandler).Methods("GET")
	s.HandleFunc("/series/edit/{urlid:[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+}", controllers.SeriesEditHandler).Methods("POST")
	s.HandleFunc("/series/add", controllers.SeriesAddFormHandler).Methods("GET")
	s.HandleFunc("/series/add", controllers.SeriesAddHandler).Methods("POST")
	s.HandleFunc("/series/{id:[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+-[a-z0-9]+}", controllers.SeriesIdHandler)
	s.HandleFunc("/series/{name}", controllers.SeriesNameHandler)
	s.HandleFunc("/series", controllers.SeriesIndexHandler)
	staticPath, _ := config.GetString("STATIC")
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(staticPath))))
	http.ListenAndServe(":3333", recovery)
}
