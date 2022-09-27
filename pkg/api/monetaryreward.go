package api

type MonetaryReward struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		AmountCents         int    `json:"amount_cents,omitempty"`
		CancellationComment string `json:"cancellation_comment,omitempty"`
		CancellationReason  string `json:"cancellation_reason,omitempty"`
		CancelledAt         string `json:"cancelled_at,omitempty"`
		Cancelled           bool   `json:"cancelled,omitempty"`
		Comment             string `json:"comment,omitempty"`
		CreatedAt           string `json:"created_at,omitempty"`
		Reason              string `json:"reason,omitempty"`
		RewardedAt          string `json:"rewarded_at,omitempty"`
		FormattedAmount     string `json:"formatted_amount,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Payments struct {
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
		} `json:"payments,omitempty"`
		RewardedBy struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"rewarded_by,omitempty"`
		Submission struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"submission,omitempty"`
		CancelledBy struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"cancelled_by,omitempty"`
	} `json:"relationships,omitempty"`
}

type MonetaryRewardResponse struct {
	Data     MonetaryReward `json:"data,omitempty"`
	Included []Identity     `json:"included,omitempty"`
}

type MonetaryRewardQuery struct {
	ID     string
	Fields struct {
		Identity       string
		MonetaryReward string
		Payment        string
		Submission     string
	}
	Include string
}

type MonetaryRewardData struct {
	Data struct {
		Type       string `json:"type,omitempty"`
		Attributes struct {
			Cancelled          bool   `json:"cancelled,omitempty"`
			CancellationReason string `json:"cancellation_reason,omitempty"`
			AmountCents        int    `json:"amount_cents,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Submission struct {
				Data struct {
					Type string `json:"type,omitempty"`
					ID   struct {
						Ref string `json:"$ref,omitempty"`
					} `json:"id,omitempty"`
				} `json:"data,omitempty"`
			} `json:"submission,omitempty"`
		} `json:"relationships,omitempty"`
	} `json:"data,omitempty"`
}
