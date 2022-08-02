package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(Writer http.ResponseWriter, Request *http.Request) {
	if err := Request.ParseForm(); err != nil {
		fmt.Fprintf(Writer, "ParseForm() error : %v", err)
		return
	}
	fmt.Fprintf(Writer, "Request Successfully Posted\n")
	Name := Request.FormValue("Name")
	Address := Request.FormValue("Address")
	fmt.Fprintf(Writer, "Name : %s\n", Name)
	fmt.Fprintf(Writer, "Address : %s", Address)
}
func helloHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path != "/hello" {
		http.Error(Writer, "404 Not Found", http.StatusNotFound)
		return
	}
	if Request.Method != "GET" {
		http.Error(Writer, "Method Not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(Writer, "Hello Guys :)")
}

func main() {
	FileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", FileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at Port: 9090\n")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
