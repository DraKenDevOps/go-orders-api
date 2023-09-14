package application

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/ok", Ok)
	router.Get("/", health)
	router.Get("/ping", basicHandler)

	return router
}

func Ok(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
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
