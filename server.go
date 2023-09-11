package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (ph *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ph.get(w, r)
	case http.MethodPost:
		ph.post(w, r)
	case http.MethodPatch, http.MethodPut:
		ph.put(w, r)
	case http.MethodDelete:
		ph.delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func idFromUrl(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		return 0, errors.New("not found")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, errors.New("not found")
	}
	return id, nil
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: Products{
			Product{"shoes", 25.00},
			Product{"Webcamp", 50.00},
			Product{"Mic", 20.00},
		},
	}
}

func main() {
	port := ":8080"
	ph := newProductHandler()
	http.Handle("/products", ph)
	http.Handle("/products/", ph)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!!")
	})

	log.Println("Server running on ", port)
	log.Fatalln(http.ListenAndServe(port, nil))

}
