package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Page struct {
	Titel   string
	Text    string
	Choices []string
}

func loadPage(title string) (*Page, error) {
	p := &Page{Titel: title}
	filename := title + ".yaml"
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(text, &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return p, nil
}

func main() {
	t, _ := loadPage("choice1.1")
	fmt.Print(t.Choices[0])
}
