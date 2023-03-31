package main

import (
	"ascii-art/runprog"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", result)
	http.ListenAndServe(":8080", nil)
}

// HOME PAGE
func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "(404) Not found: "+err.Error(), http.StatusInternalServerError)
		log.Printf("(404) Not found: %v", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

// RESULT PAGE
func result(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "(400) Bad Request", http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
		return
	}

	//get the data from the form
	input := r.Form.Get("inputText")
	if input == "" {
		http.Error(w, "(400) Bad Request: Please enter some text", http.StatusBadRequest)
		log.Println("(400) Bad Request: Please enter some text")
		return
	}

	banner := r.Form.Get("bannerType")
	result := runprog.RunProgram(input, banner)

	if result == "" {
		http.NotFound(w, r)
		log.Printf("(404) Not Found: no result found for input: %q and banner type: %q", input, banner)
		return
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "(500) Internal Server Error", http.StatusInternalServerError)
		log.Printf("(500) Internal Server Error: %v", err)
		return
	}

	type data struct {
		Rslt string
	}
	resul := data{Rslt: result}

	err = tmpl.Execute(w, resul)
	if err != nil {
		http.Error(w, "(500) Internal Server Error", http.StatusInternalServerError)
		log.Printf("(500) Internal Server Error: %v", err)
		return
	}

	log.Printf("(200) OK: result generated for input: %q and banner type: %q", input, banner)
}
