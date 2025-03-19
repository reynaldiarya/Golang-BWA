package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/rey", handler.ReyHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
