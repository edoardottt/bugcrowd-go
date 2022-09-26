package api

type Program struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		Code string `json:"code,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		CurrentBrief struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"current_brief,omitempty"`
		Organization struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"organization,omitempty"`
		Submissions struct {
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
		} `json:"submissions,omitempty"`
	} `json:"relationships,omitempty"`
}

type ProgramResponse struct {
	Data     Program        `json:"data,omitempty"`
	Included []ProgramBrief `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
	Links    Links          `json:"links,omitempty"`
}

type ProgramsResponse struct {
	Data     []Program      `json:"data,omitempty"`
	Included []ProgramBrief `json:"included,omitempty"`
	Meta     Meta           `json:"meta,omitempty"`
	Links    Links          `json:"links,omitempty"`
}
