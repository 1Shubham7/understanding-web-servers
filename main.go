package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":6000", nil)
}