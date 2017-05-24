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
	ID      []string
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
	fmt.Print(p.ID)
	return *p, nil
}

func showChoice(w http.ResponseWriter, r *http.Request) {

	p, err := loadPage("choice1")
	if err != nil {
		fmt.Print(err)
	}
	t, err := template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(r.URL)
	t.ExecuteTemplate(w, "content", p)
}

func api(w http.ResponseWriter, r *http.Request) {
	choice := r.FormValue("choice")
	fmt.Println(choice)
	http.Redirect(w, r, "/main"+"?q="+choice, http.StatusFound)
}

func main() {
	http.HandleFunc("/main", showChoice)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}
