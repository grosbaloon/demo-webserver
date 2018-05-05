package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		currentTime := time.Now()
		fmt.Fprintf(w, "Hello world. I'm %s and it's %s\n", hostname, currentTime.Format("2006-01-02 15:04:05"))
		log.Println("Request from: ", r.RemoteAddr)
	})

	log.Println("Starting webserver")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
