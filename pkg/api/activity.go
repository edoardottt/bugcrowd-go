package api

type Activity struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		CreatedAt string `json:"created_at,omitempty"`
		Key       string `json:"key,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
	} `json:"relationships,omitempty"`
}
