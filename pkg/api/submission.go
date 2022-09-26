package api

import "time"

type Submission struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		BugURL       string `json:"bug_url,omitempty"`
		CustomFields struct {
			Property1 string `json:"property1,omitempty"`
			Property2 string `json:"property2,omitempty"`
		} `json:"custom_fields,omitempty"`
		Description                         string    `json:"description,omitempty"`
		Duplicate                           bool      `json:"duplicate,omitempty"`
		ExtraInfo                           string    `json:"extra_info,omitempty"`
		HTTPRequest                         string    `json:"http_request,omitempty"`
		LastTransitionedToInformationalAt   time.Time `json:"last_transitioned_to_informational_at,omitempty"`
		LastTransitionedToNotApplicableAt   time.Time `json:"last_transitioned_to_not_applicable_at,omitempty"`
		LastTransitionedToNotReproducibleAt time.Time `json:"last_transitioned_to_not_reproducible_at,omitempty"`
		LastTransitionedToOutOfScopeAt      time.Time `json:"last_transitioned_to_out_of_scope_at,omitempty"`
		LastTransitionedToResolvedAt        time.Time `json:"last_transitioned_to_resolved_at,omitempty"`
		LastTransitionedToTriagedAt         time.Time `json:"last_transitioned_to_triaged_at,omitempty"`
		LastTransitionedToUnresolvedAt      time.Time `json:"last_transitioned_to_unresolved_at,omitempty"`
		Severity                            int       `json:"severity,omitempty"`
		RemediationAdvice                   string    `json:"remediation_advice,omitempty"`
		SubmittedAt                         time.Time `json:"submitted_at,omitempty"`
		Source                              string    `json:"source,omitempty"`
		State                               string    `json:"state,omitempty"`
		Title                               string    `json:"title,omitempty"`
		VrtID                               string    `json:"vrt_id,omitempty"`
		VrtVersion                          string    `json:"vrt_version,omitempty"`
		VulnerabilityReferences             string    `json:"vulnerability_references,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Activities struct {
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
		} `json:"activities,omitempty"`
		Assignee struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"assignee,omitempty"`
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
		ClaimTicket struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"claim_ticket,omitempty"`
		Comments struct {
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
		} `json:"comments,omitempty"`
		CvssVector struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"cvss_vector,omitempty"`
		Duplicates struct {
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
		} `json:"duplicates,omitempty"`
		DuplicateOf struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"duplicate_of,omitempty"`
		ExternalIssues struct {
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
		} `json:"external_issues,omitempty"`
		FileAttachments struct {
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
		} `json:"file_attachments,omitempty"`
		MonetaryRewards struct {
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
		} `json:"monetary_rewards,omitempty"`
		Target struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"target,omitempty"`
		Researcher struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"researcher,omitempty"`
	} `json:"relationships,omitempty"`
}

type SubmissionResponse struct {
	Data     Submission `json:"data,omitempty"`
	Included []Program  `json:"included,omitempty"`
	Meta     Meta       `json:"meta,omitempty"`
	Links    Links      `json:"links,omitempty"`
}

type SubmissionsResponse struct {
	Data     []Submission `json:"data,omitempty"`
	Included []Program    `json:"included,omitempty"`
	Meta     Meta         `json:"meta,omitempty"`
	Links    Links        `json:"links,omitempty"`
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
