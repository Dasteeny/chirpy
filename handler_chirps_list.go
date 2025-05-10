package main

import (
	"net/http"

	"github.com/dasteeny/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {
	author_id_str := r.URL.Query().Get("author_id")
	userID, err := uuid.Parse(author_id_str)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid author ID", err)
	}

	chirps := []database.Chirp{}
	if userID == uuid.Nil {
		chirps, err = cfg.db.ListChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps list", err)
			return
		}
	} else {
		chirps, err = cfg.db.ListChirpsByAuthor(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps list", err)
			return
		}
	}

	chirpsList := []Chirp{}
	for _, chirp := range chirps {
		chirpsList = append(chirpsList, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirpsList)
}
