package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type page struct {
	Title string
	Body  []byte
}

func host(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "unknown hostname")
	} else {
		fmt.Fprintf(w, name)
	}
}

func httpRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + r.URL.Path[6:])
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body[:]))
}

func random(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pseudoUUID())
}

func raw(w http.ResponseWriter, r *http.Request) {
	value, exists := os.LookupEnv("RAW_CONTENT")
	if !exists {
		value = "raw"
	}
	fmt.Fprintf(w, value)
}

func pseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func html(w http.ResponseWriter, r *http.Request) {
	p := &page{
		Title: os.Getenv("HTML_TITLE"),
		Body:  []byte(os.Getenv("HTML_BODY")),
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/host", host)
	r.HandleFunc("/raw", raw)
	r.HandleFunc("/random", random)
	r.HandleFunc("/html", html)
	r.PathPrefix("/http/").HandlerFunc(httpRequest)
	r.HandleFunc("/", raw)
	http.ListenAndServe(":8080", r)
}
