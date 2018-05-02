package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello world. I'm %s.", hostname)
		log.Println("Request from: ", r.RemoteAddr)
	})

	log.Println("Starting webserver")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
