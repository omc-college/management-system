package openapi

type roleTmpl struct {
	Entries map[string]featureEntry `json:"entries"`
}

type featureEntry struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Endpoints   map[string]endpoint `json:"endpoints"`
}

type endpoint struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type apiFeature struct {
	Description string `json:"description"`
}
