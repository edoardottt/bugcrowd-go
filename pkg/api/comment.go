package api

import "time"

type Comment struct {
	Data struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Body            string    `json:"body,omitempty"`
			VisibilityScope string    `json:"visibility_scope,omitempty"`
			CreatedAt       time.Time `json:"created_at,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Submission struct {
				Data struct {
					Type string `json:"type,omitempty"`
					ID   struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
				} `json:"data,omitempty"`
			} `json:"submission,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"data,omitempty"`
}
