package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}
	http.HandleFunc("/", loggerMiddleware(helloWorld))
	http.HandleFunc("/ping", loggerMiddleware(ping))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
		next(w, r)
	}
}
