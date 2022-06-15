package main

import (
	"log"
	"flag"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/remove/", remove)

	log.Println("Start: http://localhost:" + *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
