package opa

type RegoInput struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Token  string `json:"token"`
}

type regExpInfo struct {
	Roles   map[string]bool
	Methods map[string]bool
}

type authCacheItem struct {
	RegExp  string
	Roles   []string
	Methods []string
}

type roles struct {
	Roles map[string]role `json:"roles"`
}

type role struct {
	Name    string               `json:"name"`
	Entries map[int]featureEntry `json:"entries"`
}

type featureEntry struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Endpoints   map[int]endpoint `json:"endpoints"`
}

type endpoint struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
