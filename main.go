package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

var Version string

func getVersion() string {
	if Version == "" {
		Version = "head"
	}
	return Version
}

var env string
var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.StringVar(&env, "env", "dev", "Environment to run the server in")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		figure.Write(w, figure.NewFigure("Deploydocus", "relief", false))
		w.Write([]byte(fmt.Sprintf("\nenv: %s, v: %s\n", env, getVersion())))
	})

	fmt.Printf("Starting 🦕 on port :%d, v: %s, env: %s\n", port, getVersion(), env)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Unable to start 🦕: ", err)
	}
}
