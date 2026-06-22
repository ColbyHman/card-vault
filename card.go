package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

type Card struct {
	Id              string `json:"id"`
	Quantity        int    `json:"quantity"`
	Name            string `json:"name"`
	CollectorNumber string `json:"collectorNumber"`
	SetCode         string `json:"setCode"`
}

func importCard(w http.ResponseWriter, r *http.Request) {
	card, err := validateCardFromRequest(r)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}
	log.Info().Msg("Card Imported: " + card.Name)
	w.WriteHeader(201)
}

func listCards(w http.ResponseWriter, r *http.Request) {
	cards := []Card{
		{
			Quantity:        1,
			Name:            "Frozen in Ice",
			CollectorNumber: "0054",
			SetCode:         "MSH",
		},
		{
			Quantity:        1,
			Name:            "Psychic Whorl",
			CollectorNumber: "0105",
			SetCode:         "BLB",
		},
	}

	if err := json.NewEncoder(w).Encode(cards); err != nil {
		log.Error().Msg("Unexpected error: " + err.Error())
		w.WriteHeader(500)
	}
}

func getCard(w http.ResponseWriter, r *http.Request) {
	cards := []Card{
		{
			Id:              "123",
			Quantity:        1,
			Name:            "Frozen in Ice",
			CollectorNumber: "0054",
			SetCode:         "MSH",
		},
		{
			Id:              "456",
			Quantity:        1,
			Name:            "Psychic Whorl",
			CollectorNumber: "0105",
			SetCode:         "BLB",
		},
	}

	id := chi.URLParam(r, "id")

	for _, card := range cards {
		if card.Id == id {
			if err := json.NewEncoder(w).Encode(card); err != nil {
				log.Error().Msg("Unexpected error: " + err.Error())
				w.WriteHeader(500)
			}
		}
	}

}

func validateCardFromRequest(r *http.Request) (Card, error) {
	defer r.Body.Close()

	var req Card
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Msg("Request contained invalid JSON.")
		return Card{}, err
	}

	return req, nil
}
