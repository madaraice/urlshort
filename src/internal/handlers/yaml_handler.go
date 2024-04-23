package handlers

import (
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
)

func YAMLHandler(yamlPath string, fallback http.Handler) (http.HandlerFunc, error) {
	pathToUrls, err := parseYAML(yamlPath)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(pathToUrls)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yamlPath string) ([]pathToUrl, error) {
	yamlContent, err := os.ReadFile(yamlPath)
	if err != nil && err != io.EOF {
		return nil, err
	}

	var pathToUrls []pathToUrl
	err = yaml.Unmarshal(yamlContent, &pathToUrls)
	return pathToUrls, err
}
