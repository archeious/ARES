package controllers

import (
	"html/template"
	"io"
	"log"
)

func substring(s string, i, c int) string {
	if c > len(s) {
		c = len(s)
	}
	return s[i:c]
}

func render(w io.Writer, c, v string, args interface{}) {
	t, err := template.New("layout.tmpl").Funcs(template.FuncMap{
		"substring": substring,
	}).ParseFiles("templates/layout.tmpl", "templates/"+c+"/"+v+".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, args); err != nil {
		log.Fatal(err)
	}
}
