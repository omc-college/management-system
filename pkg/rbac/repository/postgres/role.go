package postgres

type role struct {
	ID      int                  `json:"id"`
	Name    string               `json:"name"`
	Entries map[int]featureEntry `json:"entries"`
}

type featureEntry struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Endpoints map[int]endpoint `json:"endpoints"`
}

type endpoint struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
