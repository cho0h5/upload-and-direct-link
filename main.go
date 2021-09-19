package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	var i interface{}
	t.Execute(w, i)
	log.Println(r.Method, r.URL.String(), r.Proto)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "upload")
	log.Println(r.Method, r.URL.String(), r.Proto)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8082", nil)
}
