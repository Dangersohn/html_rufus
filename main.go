package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Titel string
	Text  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Titel: title, Text: text}, nil
}

func showChoice(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("choice1.1")
	t, _ := template.ParseGlob("template/*.html")
	t.ExecuteTemplate(w, string(p.Text), r.RequestURI)
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.FormValue("choice"))
}

func main() {
	http.HandleFunc("/test", showChoice)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}
