package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

var (
	Version   string = "head"
	BuildDate string = "-"
	env       string
	port      int
)

func init() {
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.StringVar(&env, "env", "dev", "Environment to run the server in")
	flag.Parse()
}

func stats() string {
	return fmt.Sprintf("env: %s, v: %s, build date: %s\n", env, Version, BuildDate)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		figure.Write(w, figure.NewFigure("Deploydocus", "relief", false))
		w.Write([]byte(fmt.Sprintf("\n%s", stats())))
	})

	fmt.Printf("Starting 🦕 on port :%d, %s\n", port, stats())
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Unable to start 🦕: ", err)
	}
}
