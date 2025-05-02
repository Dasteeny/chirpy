package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	uuidChirpID, err := uuid.Parse(chirpID)
	if err != nil {
		http.Error(w, "invalid chirp ID", http.StatusBadRequest)
		return
	}
	chirp, err := cfg.db.GetChirp(r.Context(), uuidChirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve a chirp", err)
		return
	}
	if chirp.ID == uuid.Nil {
		respondWithJSON(w, http.StatusNotFound, nil)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
