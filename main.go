package main

// http server

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "port to listen on")
	showHeader := flag.Bool("show-header", true, "show header")

	flag.Parse()

	log.Println("Args:", flag.Args())

	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, r.URL.Path)

		responseBody := ""

		if *showHeader {
			log.Println("=====================header=====================")
			responseBody += "=====================header=====================" + "\n"
			for k, v := range r.Header {
				log.Println("==>", k, ": ", v)
				responseBody += k + ": " + v[0] + "\n"
			}
			responseBody += "====================end header===================" + "\n"
			log.Println("====================end header===================")
		}

		query := r.URL.Query()
		log.Println("=====================query=====================")
		responseBody += "=====================query=====================" + "\n"
		for k, v := range query {
			log.Println("==>", k, ": ", v)
			responseBody += k + ": " + v[0] + "\n"
		}
		responseBody += "====================end query===================" + "\n"
		log.Println("===================end query======================")

		// Read body to string
		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Printf("error reading body: %v", err)
			http.Error(w, "error reading body", http.StatusInternalServerError)
			return
		}

		log.Println("=====================body=====================")
		responseBody += "=====================body=====================" + "\n"
		log.Println("==>", string(body))
		responseBody += string(body) + "\n"
		responseBody += "====================end body===================" + "\n"
		log.Println("===================end body=====================")

		// Write response
		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write([]byte(responseBody))
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
