package main

import (
	"log"
	"flag"
	"net/http"
	"io/ioutil"
	"html/template"

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
	port := flag.String("p", "8080", "port")
	flag.Parse()

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)

	log.Println("Start: http://localhost:" + *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
