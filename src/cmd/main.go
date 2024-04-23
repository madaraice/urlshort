package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"urlshort/src/internal/handlers"
)

var (
	yamlPath = flag.String("yaml", "", "Path to YAML file")
	jsonPath = flag.String("json", "", "Path to JSON file")
)

func main() {
	flag.Parse()
	if *yamlPath == "" || *jsonPath == "" {
		fmt.Println("Must specify -yaml and -json")
		os.Exit(1)
	}

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	yamlHandler, err := handlers.YAMLHandler(*yamlPath, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := handlers.JSONHandler(*jsonPath, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")

	err = http.ListenAndServe(":8080", jsonHandler)
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
