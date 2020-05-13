package rbac

type Endpoint struct {
	Id     int    `json:"id"`
	Method string `json:"method"`
	Name   string `json:"name"`
	Path   string `json:"path"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type FeatureEntry struct {
	Description string     `json:"description"`
	Endpoints   []Endpoint `json:"endpoints"`
	Id          int        `json:"id"`
	Name        string     `json:"name"`
}

type Role struct {
	Entries []FeatureEntry `json:"entries"`
	Id      int            `json:"id"`
	Name    string         `json:"name"`
}

type RoleTemplate struct {
	Entries []FeatureEntry `json:"entries"`
}
