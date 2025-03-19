package handler

import (
	"fmt"
	"golangweb/entity"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to the home page"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ReyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title": "Learning Golang Web",
	// 	"content": "Hello from the template",		
	// }

	data:= entity.Product{ID: 1, Name: "Mobil", Price: 100000, Stock: 10}

	err = tmpl.Execute(w, data)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Product page: %d", idNumb)
}
