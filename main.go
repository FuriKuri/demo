package main

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: os.Getenv("HTML_TITLE"),
		Body:  []byte(os.Getenv("HTML_BODY")),
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
