package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var Version string

func getVersion() string {
	if Version == "" {
		Version = "unknown"
	}
	return Version
}

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Running 🦕 version: %s\n", getVersion())))
	})

	fmt.Printf("Starting 🦕 on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Unable to start 🦕: ", err)
	}
}
