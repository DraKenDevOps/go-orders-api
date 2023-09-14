package main

import (
	"fmt"
	"net/http"
	// "github.com/go-chi/chi/v5"
)

func main() {
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: http.HandlerFunc(basicHandler),
	}

	error := server.ListenAndServe()

	if error != nil {
		fmt.Printf("Hello Microservice Moon 🚀")
	} else {
		fmt.Print(error)
	}
}

func basicHandler(res http.ResponseWriter, req *http.Request) {
	// fmt.Print(req)
	res.Write([]byte("Hello Microservice Moon 🚀"))
}
