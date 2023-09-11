package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJSON(w, code, map[string]string{"error": msg})
}

func responseWithJSON(w http.ResponseWriter, code int, data interface{}) {
	// response, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	// w.Write(response)
	json.NewEncoder(w).Encode(data)
}
