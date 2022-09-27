package api

type Target struct {
	Data []struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Name     string `json:"name,omitempty"`
			Category string `json:"category,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
		} `json:"relationships,omitempty"`
	} `json:"data,omitempty"`
	Included []struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Name string `json:"name,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Programs struct {
				Data []struct {
					ID struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"data,omitempty"`
				Links struct {
					Related struct {
						Href string `json:"href,omitempty"`
						Meta struct {
							Count     int `json:"count,omitempty"`
							TotalHits int `json:"total_hits,omitempty"`
						} `json:"meta,omitempty"`
					} `json:"related,omitempty"`
				} `json:"links,omitempty"`
			} `json:"programs,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"included,omitempty"`
	Meta struct {
		TotalHits int `json:"total_hits,omitempty"`
		Count     int `json:"count,omitempty"`
	} `json:"meta,omitempty"`
	Links struct {
		Self     string `json:"self,omitempty"`
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
	} `json:"links,omitempty"`
}

type TargetQuery struct {
	Id     string
	Fields struct {
		Organization string
		Target       string
	}
	Filter struct {
		OrganizationId string
		TargetId       string
	}
	Include string
	Sort    string
}

type TargetsResponse struct {
	Data     []Target       `json:"data,omitempty"`
	Included []Organization `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
	Links    Links          `json:"links,omitempty"`
}
