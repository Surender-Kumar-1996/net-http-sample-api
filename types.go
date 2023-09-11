package main

import "sync"

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Products []Product

type productHandler struct {
	sync.Mutex
	products Products
}
