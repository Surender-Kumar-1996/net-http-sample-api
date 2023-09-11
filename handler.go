package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func (ph *productHandler) get(w http.ResponseWriter, r *http.Request) {
	defer ph.Unlock()
	ph.Lock()
	id, err := idFromUrl(r)
	if err != nil {
		responseWithJSON(w, http.StatusOK, ph.products)
		return
	}
	if id >= len(ph.products) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	responseWithJSON(w, http.StatusOK, ph.products[id])

}
func (ph *productHandler) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}
	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer ph.Unlock()
	ph.Lock()
	ph.products = append(ph.products, product)
	responseWithJSON(w, http.StatusOK, product)
}
func (ph *productHandler) put(w http.ResponseWriter, r *http.Request) {
	id, err := idFromUrl(r)
	if err != nil {
		responseWithJSON(w, http.StatusNotFound, err.Error())
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}
	defer ph.Unlock()
	ph.Lock()
	if id >= len(ph.products) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	var product Product
	err = json.Unmarshal(body, &product)
	if product.Name != "" {
		ph.products[id].Name = product.Name
	}
	if product.Price != 0.0 {
		ph.products[id].Price = product.Price
	}

	responseWithJSON(w, http.StatusOK, ph.products[id])
}
func (ph *productHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := idFromUrl(r)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	defer ph.Unlock()
	ph.Lock()
	if id >= len(ph.products) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	if id < len(ph.products)-1 {
		ph.products[len(ph.products)-1], ph.products[id] = ph.products[id], ph.products[len(ph.products)-1]
	}
	ph.products = ph.products[:len(ph.products)-1]
	responseWithJSON(w, http.StatusOK, ph.products)
}
