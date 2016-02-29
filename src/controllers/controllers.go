package controllers

import (
	"html/template"
	"io"
	"log"
)

func render(w io.Writer, c, v string, args interface{}) {
	t, err := template.ParseFiles("templates/layout.tmpl", "templates/"+c+"/"+v+".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, args); err != nil {
		log.Fatal(err)
	}
}
