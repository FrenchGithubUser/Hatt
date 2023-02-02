package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/itemsList", getItemsList)

	http.ListenAndServe(":8081", nil)
}
