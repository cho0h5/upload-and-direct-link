package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		log.Println(file.Name())
	}

	t, _ := template.ParseFiles("index.html")
	var i interface{}
	t.Execute(w, i)

	log.Println(r.Method, r.URL.String(), r.Proto)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, fileHeader, _ := r.FormFile("fileName")
	defer file.Close()
	
	fileByte, _ := ioutil.ReadAll(file)
	ioutil.WriteFile(fileHeader.Filename, fileByte, 0644)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	log.Println(r.Method, r.URL.String(), r.Proto)
	log.Println("Uploaded File:", fileHeader.Filename)
	log.Println("File Size:", fileHeader.Size)
}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8082", nil)
}
