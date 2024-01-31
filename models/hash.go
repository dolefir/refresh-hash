package models

import "time"

// Hash defines a data model.
type Hash struct {
	ID       string    `json:"uuid"`
	Datatime time.Time `json:"datatime"`
}
