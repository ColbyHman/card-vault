package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	displayBanner()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var r *chi.Mux = initServer()

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error().Msg("An error occured: " + err.Error())
	}

}
