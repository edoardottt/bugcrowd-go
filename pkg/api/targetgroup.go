package api

type TargetGroup struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		InScope     bool   `json:"in_scope,omitempty"`
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
		ProgramBrief struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"program_brief,omitempty"`
		RewardRange struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"reward_range,omitempty"`
	} `json:"relationships,omitempty"`
}
