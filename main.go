package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Addres : %s\n", address)
}

func hyHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hy" {
		http.Error(w, "Page not found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hy")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hy", hyHandler)

	fmt.Println("Server stated at port 8000")
	if error := http.ListenAndServe(":8000", nil); error != nil {
		log.Fatal(error)
	}
}
