package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("I love GeeksforGeeks")
        
		d, _ := ioutil.ReadAll(r.Body)
        
		log.Printf("Data : %s\n", d)

		fmt.Fprintf(rw, "\nHello There %s", d)
	})
	http.ListenAndServe(":6000", nil)
}

