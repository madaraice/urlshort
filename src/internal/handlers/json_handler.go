package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func JSONHandler(jsonPath string, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := parseJSON(jsonPath)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(pathToUrls)
	return MapHandler(pathMap, fallback), nil
}

func parseJSON(jsonPath string) ([]pathToUrl, error) {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil && err != io.EOF {
		return nil, err
	}

	var pathToUrls []pathToUrl
	err = json.Unmarshal(jsonContent, &pathToUrls)
	return pathToUrls, err
}
