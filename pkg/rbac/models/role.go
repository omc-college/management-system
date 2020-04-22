package models

type Role struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Entries []FeatureEntry `json:"entries"`
}

type FeatureEntry struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
