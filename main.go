package main

import (
	"fmt"
	"log"
	"net/http"

	"chipskein/htmx-experimental/controllers"
)

func main() {
	var PORT = 8081
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/animes", controllers.HandleRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
