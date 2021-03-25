package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %.2f\n", item, price)
	}
}

func main() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
