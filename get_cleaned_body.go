package main

import "strings"

func getCleanedBody(chirp string) string {
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(chirp, " ")

	for i, word := range words {
		if _, exist := badWords[strings.ToLower(word)]; exist {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
