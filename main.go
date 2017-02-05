package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", Index)
	mux.HandleFunc("/1", One)
	mux.HandleFunc("/2", Two)
	// mux.Handle("/static",http.FileServer(http.Dir("./static")))

	s := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Listening...")
	log.Fatal(s.ListenAndServe())
}
