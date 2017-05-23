package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func showChoice(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("template/*.html")
	t.ExecuteTemplate(w, "test", r.RequestURI)
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.FormValue("choice"))
}

func main() {
	http.HandleFunc("/test", showChoice)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}
