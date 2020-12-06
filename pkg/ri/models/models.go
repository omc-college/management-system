package models

import "time"

type Room struct {
	ID			int		  `json:"id" db:"id"`
	Room		string	  `json:"room" db:"room"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
}

type Resources struct {
	ID                  int       `json:"id" db:"id"`
	ResourceName        string    `json:"resource_name" db:"resource_name"`
	ResourceDescription string    `json:"resource_description" db:"resource_description"`
	ModifiedAt          time.Time `json:"modified_at" db:"modified_at"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}