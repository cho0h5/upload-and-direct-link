package main

import (
	// "log"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "upload")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8082", nil)
}
