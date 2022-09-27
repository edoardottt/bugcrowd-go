package api

type Payment struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		AmountCents int    `json:"amount_cents,omitempty"`
		CreatedAt   string `json:"created_at,omitempty"`
		RemittedAt  string `json:"remitted_at,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		MonetaryReward struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"monetary_reward,omitempty"`
		RemittedBy struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"remitted_by,omitempty"`
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
