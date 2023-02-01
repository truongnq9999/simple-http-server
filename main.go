package main

// http server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	port := flag.String("port", "8080", "port to listen on")
	//showHeader := flag.Bool("show-header", true, "show header")

	flag.Parse()

	log.Println("Args:", flag.Args())

	handler := func(w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("REQUEST:\n%s", string(reqDump))

		// Write response
		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write(reqDump)
		if err != nil {
			return
		}

		if err != nil {
			log.Printf("error: %v", err)
			return
		}
	}

	addr := ":" + *port

	log.Printf("Starting server at %v", addr)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
