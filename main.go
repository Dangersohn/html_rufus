package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"regexp"

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
	text, err := ioutil.ReadFile("choices/" + filename)
	if err != nil {
		text, err = ioutil.ReadFile("choices/intro.yaml")
		if err != nil {
			return *p, err
		}
	}
	err = yaml.Unmarshal(text, &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//	fmt.Print(p.ID)
	return *p, nil
}

func showChoice(w http.ResponseWriter, r *http.Request) {
	var re = regexp.MustCompile(`[^=]+$`)
	choiceQuery := r.URL.RawQuery // Schneidet das Query aus
	match := re.FindAllString(choiceQuery, 1)
	p, err := loadPage(match[0])
	if err != nil {
		fmt.Print(err)
	}
	t, err := template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(r.URL)

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
