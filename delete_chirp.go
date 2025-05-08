package main

import (
	"net/http"

	"github.com/dasteeny/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	uuidChirpID, err := uuid.Parse(chirpID)
	if err != nil {
		http.Error(w, "invalid chirp ID", http.StatusBadRequest)
		return
	}

	jwtToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(jwtToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), uuidChirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve a chirp", err)
		return
	}
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Not allowed", nil)
		return
	}

	err = cfg.db.DeleteChirp(r.Context(), uuidChirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
