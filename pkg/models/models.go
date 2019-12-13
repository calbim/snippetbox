package models

import (
	"errors"
	"time"
)

//ErrRecordNotFound is the error returned when no matching record is found
var ErrRecordNotFound = errors.New("models: no matching record found")

//Snippet represents a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
