package postgres

import "database/sql"

type role struct {
	ID      int                            `json:"id"`
	Name    string                         `json:"name"`
	Entries map[sql.NullInt64]featureEntry `json:"entries"`
}

type featureEntry struct {
	ID          sql.NullInt64              `json:"id"`
	Name        sql.NullString             `json:"name"`
	Description sql.NullString             `json:"description"`
	Endpoints   map[sql.NullInt64]endpoint `json:"endpoints"`
}

type endpoint struct {
	ID     sql.NullInt64  `json:"id"`
	Name   sql.NullString `json:"name"`
	Path   sql.NullString `json:"path"`
	Method sql.NullString `json:"method"`
}

type roleTmpl struct {
	Entries map[sql.NullInt64]featureEntry `json:"entries"`
}
