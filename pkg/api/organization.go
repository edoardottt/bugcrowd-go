package api

type Organization struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		Name string `json:"name,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Targets struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
					Meta struct {
						Count     int `json:"count,omitempty"`
						TotalHits int `json:"total_hits,omitempty"`
					} `json:"meta,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data []struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"targets,omitempty"`
		Programs struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
					Meta struct {
						Count     int `json:"count,omitempty"`
						TotalHits int `json:"total_hits,omitempty"`
					} `json:"meta,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data []struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"programs,omitempty"`
	} `json:"relationships,omitempty"`
}

type OrganizationQuery struct {
	ID     string
	Fields struct {
		Organization string
		Program      string
		Target       string
	}
	Include string
	Sort    string
}

type OrganizationResponse struct {
	Data     Organization `json:"data,omitempty"`
	Included []Program    `json:"included,omitempty"`
	Meta     Meta         `json:"meta,omitempty"`
	Links    Links        `json:"links,omitempty"`
}

type OrganizationsResponse struct {
	Data     []Organization `json:"data,omitempty"`
	Included []Program      `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
	Links    Links          `json:"links,omitempty"`
}
