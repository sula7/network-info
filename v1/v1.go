package v1

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Start(r *chi.Mux, bindPort string) error {
	log.Printf("Listen on %v", bindPort)
	return http.ListenAndServe(bindPort, r)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
