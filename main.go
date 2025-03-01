package main

import (
	"flag"
	"fmt"
	"html/template"
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

const pageTmpl = `
<html>
	<head>
		<meta name="color-scheme" content="dark">
	    <meta charset="UTF-8">
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🦕</text></svg>">
	</head>
	<body>
		<pre style="word-wrap: break-word;">{{.Logo}}</pre>
		<pre style="word-wrap: break-word; white-space: pre-wrap;">{{.Stats}}</pre>
		Oh no, it's a bug!
	</body>
</html>`

type pageTmplData struct {
	Logo  string
	Stats string
}

func main() {
	tmpl, err := template.New("page").Parse(pageTmpl)
	if err != nil {
		log.Fatal("Failed to parse page template: ", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err = tmpl.Execute(w, pageTmplData{
			Logo:  figure.NewFigure("Deploydocus", "relief", false).String(),
			Stats: stats(),
		})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	fmt.Printf("Starting 🦕 on port :%d, %s\n", port, stats())
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Unable to start 🦕: ", err)
	}
}
