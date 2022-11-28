package main

import (
	http1 "Project3/Placement/http"
	"net/http"
)

func main() {

	h1 := http1.New()
	h2 := http1.New()

	http.HandleFunc("/student", h1.handler)
	http.HandleFunc("/company", h2.handler)
	http.ListenAndServe(":8080", nil)
}
