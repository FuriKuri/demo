package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func raw(w http.ResponseWriter, r *http.Request) {
	value, exists := os.LookupEnv("RAW_CONTENT")
	if !exists {
		value = pseudo_uuid()
	}
	fmt.Fprintf(w, value)
}

func pseudo_uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
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
	http.HandleFunc("/raw", raw)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
