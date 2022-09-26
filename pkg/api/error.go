package api

import "fmt"

type Error struct {
	Title  string `json:"title,omitempty"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
	} `json:"source,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Detail)
}
