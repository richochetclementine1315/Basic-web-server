package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name") //in name varible name from form is stored
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s\n", address)
}

// API request is r and response is w
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "HELLO!!")

}

func main() {
	//Telling Golang to basically checkout the static dir and then looks at the index.html file.same as php node js etc.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)          //handle function in http package basically handles the root route that is the "/"route(pattern string) and send it to the fileServer(http.handler)
	http.HandleFunc("/form", formHandler) //same as http.handle but used when a pointer of request type is there.
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
