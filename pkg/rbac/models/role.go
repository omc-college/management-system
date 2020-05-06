package models

type Role struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Entries []FeatureEntry `json:"entries"`
}

type FeatureEntry struct {
	ID          int        `json:"id,omitempty" yaml:"entries,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Endpoints   []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ID     int    `json:"id,omitempty" yaml:"entries,omitempty"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type RoleTmpl struct {
	Entries []FeatureEntry `json:"entries"`
}
