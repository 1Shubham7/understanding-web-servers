package main

import (
	"net/http"
	"github.com/1shubham7/understanding-web-servers/handlers"
	"log"
	"os"
	// "io"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	homehandler := handlers.NewHome(l) //you can just say it hh

	servemux := http.NewServeMux() //you can also call this sm
	servemux.Handle("/", homehandler)

	http.ListenAndServe(":6000", servemux)
	// second parameter is for http handler. if we say "nil" in second parameter, 
	// the server will take the default http handler, here we specified the http handler



	
}

