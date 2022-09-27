package api

import "time"

type RewardRange struct {
	Data struct {
		ID struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			AmountCents         int         `json:"amount_cents,omitempty"`
			CancellationComment interface{} `json:"cancellation_comment,omitempty"`
			CancellationReason  string      `json:"cancellation_reason,omitempty"`
			CancelledAt         time.Time   `json:"cancelled_at,omitempty"`
			Comment             interface{} `json:"comment,omitempty"`
			CreatedAt           time.Time   `json:"created_at,omitempty"`
			Reason              string      `json:"reason,omitempty"`
			RewardedAt          time.Time   `json:"rewarded_at,omitempty"`
			FormattedAmount     string      `json:"formatted_amount,omitempty"`
			Cancelled           bool        `json:"cancelled,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Submission struct {
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
			} `json:"submission,omitempty"`
			RewardedBy struct {
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
			} `json:"rewarded_by,omitempty"`
			CancelledBy struct {
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
			} `json:"cancelled_by,omitempty"`
			Payments struct {
				Ref string `json:"$ref,omitempty"`
			} `json:"payments,omitempty"`
		} `json:"relationships,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"data,omitempty"`
}
