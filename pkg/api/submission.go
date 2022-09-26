package api

type SubmissionResponse struct {
	Data []struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			State     string `json:"state,omitempty"`
			Duplicate bool   `json:"duplicate,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
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

type SubmissionQuery struct {
	Id     string
	Fields struct {
		Activity       string
		ClaimTicket    string
		Comment        string
		CvssVector     string
		ExternalIssue  string
		FileAttachment string
		Identity       string
		MonetaryReward string
		Organization   string
		Payment        string
		Program        string
		ProgramBrief   string
		Submission     string
		Target         string
		TargetGroup    string
	}
	Filters struct {
		Assignee   string
		BlockedBy  string
		Duplicate  bool
		Payments   string
		Points     string
		Program    string
		Researcher string
		Retest     string
		Search     string
		Severity   string
		Source     string
		State      string
		Submitted  string
		Target     string
		TargetType string
		Updated    string
		Vrt        string
		Sort       string
	}
	Include string
}

type SubmissionData struct {
	Data struct {
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Title       string `json:"title,omitempty"`
			Description string `json:"description,omitempty"`
			Severity    int    `json:"severity,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Program struct {
				Data struct {
					Type string `json:"type,omitempty"`
					ID   struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
				} `json:"data,omitempty"`
			} `json:"program,omitempty"`
		} `json:"relationships,omitempty"`
	} `json:"data,omitempty"`
}
