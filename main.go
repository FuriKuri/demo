package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func echo(w http.ResponseWriter, r *http.Request) {
	var request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request += url
	request += fmt.Sprintf("\nHost: %v", r.Host)
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request += fmt.Sprintf("\n%v: %v", name, h)
		}
	}

	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	request += fmt.Sprintf("\n\nBody: %v", string(bodyBuffer))

	fmt.Fprintf(w, "%s", request)
}

func delay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	delay, _ := strconv.Atoi(vars["time"])
	fmt.Println(delay)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Fprintf(w, "delay: "+vars["time"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/host", host)
	r.HandleFunc("/echo", echo)
	r.HandleFunc("/raw", raw)
	r.HandleFunc("/random", random)
	r.HandleFunc("/delay/{time:[0-9]+}", delay)
	r.HandleFunc("/html", html)
	r.PathPrefix("/http/").HandlerFunc(httpRequest)
	r.HandleFunc("/", raw)
	fmt.Println("Listen on port 8080...")
	http.ListenAndServe(":8080", r)
}
