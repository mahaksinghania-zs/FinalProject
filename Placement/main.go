package main

import (
	"net/http"
)

func main() {

	//h1 := http1.New()
	//h2 := http1.New()
	//
	//http.HandleFunc("/student", h1.Handler)
	//http.HandleFunc("/company", h2.Handler)
	http.ListenAndServe(":8080", nil)
}
