package api

import "time"

type ClaimTicket struct {
	Data struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			ClaimedAt interface{} `json:"claimed_at,omitempty"`
			ExpiresAt time.Time   `json:"expires_at,omitempty"`
			Status    string      `json:"status,omitempty"`
			Token     string      `json:"token,omitempty"`
			ClaimURL  string      `json:"claim_url,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Submissions struct {
				Data []struct {
					Type string `json:"type,omitempty"`
					ID   struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
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
			} `json:"submissions,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"data,omitempty"`
}
