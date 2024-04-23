package handlers

type pathToUrl struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func buildMap(pathToUrls []pathToUrl) map[string]string {
	pathMap := make(map[string]string, len(pathToUrls))
	for _, pu := range pathToUrls {
		pathMap[pu.Path] = pu.Url
	}

	return pathMap
}
