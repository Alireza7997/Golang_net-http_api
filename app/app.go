package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alireza/routers"
)

func App() {
	routers.SetUpRouters()

	fmt.Printf("Starting Server at Port: 9090\n")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
