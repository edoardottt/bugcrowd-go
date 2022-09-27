package api

import "time"

type Event struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		CreatedAt time.Time `json:"created_at,omitempty"`
		Key       string    `json:"key,omitempty"`
		Data      struct {
			BlockedBy string `json:"blocked_by,omitempty"`
		} `json:"data,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Actor struct {
			Data struct {
				Type string `json:"type,omitempty"`
				ID   string `json:"id,omitempty"`
			} `json:"data,omitempty"`
		} `json:"actor,omitempty"`
		Resource struct {
			Data struct {
				Type string `json:"type,omitempty"`
				ID   string `json:"id,omitempty"`
			} `json:"data,omitempty"`
		} `json:"resource,omitempty"`
	} `json:"relationships,omitempty"`
}
