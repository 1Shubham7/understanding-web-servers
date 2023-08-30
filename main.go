package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("I love GeeksforGeeks")
})
	http.ListenAndServe(":6000", nil)
}

