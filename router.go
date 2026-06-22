package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

type HealthStatus struct {
	Status string `json:"status"`
}

func initServer() *chi.Mux {
	log.Debug().Msg("Creating Card Vault Router...")

	var r = chi.NewRouter()

	r.Route("/cards", func(r chi.Router) {
		r.Post("/", importCard)
		r.Get("/", listCards)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getCard)
		})
	})

	r.Route("/health", func(r chi.Router) {
		r.Get("/", serverHealth)
	})

	return r
}

func serverHealth(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Server Health OK")
	status := HealthStatus{
		Status: "ok",
	}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Error().Msg("Unexpected error: " + err.Error())
		w.WriteHeader(500)
	}
}
