package routers

import (
	"net/http"

	"github.com/alireza/handlers"
)

func SetUpRouters() {
	FileServer := http.FileServer(http.Dir("templates/"))
	http.Handle("/", FileServer)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)
}
