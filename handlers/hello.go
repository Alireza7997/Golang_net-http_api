package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(Writer http.ResponseWriter, Request *http.Request) {
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
