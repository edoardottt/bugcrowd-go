package api

type ProgramBrief struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		CoordinatedDisclosure bool   `json:"coordinated_disclosure,omitempty"`
		CreatedAt             string `json:"created_at,omitempty"`
		Demo                  bool   `json:"demo,omitempty"`
		Description           string `json:"description,omitempty"`
		Tagline               string `json:"tagline,omitempty"`
		RewardsOverview       string `json:"rewards_overview,omitempty"`
		TargetsOverview       string `json:"targets_overview,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Program struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"program,omitempty"`
		TargetGroups struct {
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
		} `json:"target_groups,omitempty"`
	} `json:"relationships,omitempty"`
}
