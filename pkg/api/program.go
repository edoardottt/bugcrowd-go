package api

type Program struct {
	Data []struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Code string `json:"code,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Organization struct {
				Data struct {
					ID struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"data,omitempty"`
				Links struct {
					Related struct {
						Href string `json:"href,omitempty"`
					} `json:"related,omitempty"`
				} `json:"links,omitempty"`
			} `json:"organization,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"data,omitempty"`
	Included []struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Demo bool `json:"demo,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Program struct {
				Data struct {
					ID struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"data,omitempty"`
				Links struct {
					Related struct {
						Href string `json:"href,omitempty"`
					} `json:"related,omitempty"`
				} `json:"links,omitempty"`
			} `json:"program,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"included,omitempty"`
	Meta struct {
		Count     int `json:"count,omitempty"`
		TotalHits int `json:"total_hits,omitempty"`
	} `json:"meta,omitempty"`
	Links struct {
		Self     string `json:"self,omitempty"`
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
	} `json:"links,omitempty"`
}
