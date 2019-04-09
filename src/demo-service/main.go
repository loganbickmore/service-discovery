package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logRequest(r *http.Request) {
	log.Printf("Request recieved from host \"%v\" for \"%v\"", r.Host, r.URL)
}

func hello(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

func ping(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "pong")
}

func fail(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func startServer() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/fail", fail)

	log.Printf("Starting server on port: %v", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
