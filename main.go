package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", health)
	router.Get("/ping", basicHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	error := server.ListenAndServe()

	if error != nil {
		fmt.Printf("Hello Microservice Moon 🚀")
	} else {
		fmt.Print(error)
	}
}

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func health(res http.ResponseWriter, req *http.Request) {
	// fmt.Println(uptime())
	res.Write([]byte("Microservice to Moon 🚀"))
}

func basicHandler(res http.ResponseWriter, req *http.Request) {
	// fmt.Print(req)
	res.Write([]byte("qwertyuiopasdfghjklxcvbnm 🚀"))
}
