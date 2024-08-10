package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marsal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Reponding with 5XX error:", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})

}
