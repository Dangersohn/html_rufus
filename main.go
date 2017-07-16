package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"

	"github.com/julienschmidt/httprouter"
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

func showChoic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, err := loadPage(ps.ByName("choice"))
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

func api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	choice := r.FormValue("choice")
	fmt.Println(choice)
	http.Redirect(w, r, "/main"+choice, http.StatusFound)
}

func main() {
	router := httprouter.New()
	router.GET("/api", api)
	router.GET("/main/:choice", showChoic)
	router.NotFound = http.FileServer(http.Dir("/static/*"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
