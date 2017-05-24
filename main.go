package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Page struct {
	Titel   string
	Text    string
	Choices []string
}

func loadPage(title string) (Page, error) {
	p := &Page{Titel: title}
	filename := title + ".yaml"
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return *p, err
	}
	err = yaml.Unmarshal(text, &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return *p, nil
}

func showChoice(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("choice1.1")
	if err != nil {
		fmt.Print(err)
	}
	t, err := template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Print(err)
	}
	t.ExecuteTemplate(w, "content", p)
}

func api(w http.ResponseWriter, r *http.Request) {
	choice := r.FormValue("choice")
	if choice == "1" {
		fmt.Fprint(w, "Hat aus 1 gewählt")
	}
	if choice == "2" {
		fmt.Fprint(w, "Hat aus 2 gewählt")
	}

}

func main() {
	http.HandleFunc("/test", showChoice)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}
