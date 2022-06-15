package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

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