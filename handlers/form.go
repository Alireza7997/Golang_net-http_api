package handlers

import (
	"fmt"
	"net/http"
)

func FormHandler(Writer http.ResponseWriter, Request *http.Request) {
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
