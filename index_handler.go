package main

import (
	"log"
	"net/http"
	"html/template"
	"io/ioutil"

	_ "embed"
)

//go:embed index.html
var indexPage string

func index(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(".")
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	t := template.New("index")
	t, _ = t.Parse(indexPage)
	t.Execute(w, fileNames)

	log.Println(r.Method, r.URL.String(), r.Proto)
}