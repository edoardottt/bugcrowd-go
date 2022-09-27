package api

type Identity struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
		Staff bool   `json:"staff,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
	} `json:"relationships,omitempty"`
}
