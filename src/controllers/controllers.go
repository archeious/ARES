package controllers

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

type key int

const MyKey key = 0

var Store = sessions.NewCookieStore([]byte("something-very-secret"))

func substring(s string, i, c int) string {
	if c > len(s) {
		c = len(s)
	}
	return s[i:c]
}

func UserAuthMiddleWare(n http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "UAMW Get Error:"+err.Error(), 500)
			return
		}
		user, ok := session.Values["user"]
		if ok {
			context.Set(r, "user", user)
		}

		n.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

func render(w http.ResponseWriter, r *http.Request, c, v string, args map[string]interface{}) {
	t, err := template.New("layout.tmpl").Funcs(template.FuncMap{
		"substring": substring,
	}).ParseFiles("templates/layout.tmpl", "templates/"+c+"/"+v+".tmpl")
	if err != nil {
		log.Fatal(err)
	}
	if args == nil {
		args = make(map[string]interface{})
	}
	args["user"] = context.Get(r, "user")
	if err := t.Execute(w, args); err != nil {
		log.Fatal(err)
	}
}
