package api

type CVSSVector struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		AttackComplexity      string `json:"attack_complexity,omitempty"`
		AttackVector          string `json:"attack_vector,omitempty"`
		AuthorizationScope    string `json:"authorization_scope,omitempty"`
		AvailabilityImpact    string `json:"availability_impact,omitempty"`
		ConfidentialityImpact string `json:"confidentiality_impact,omitempty"`
		IntegrityImpact       string `json:"integrity_impact,omitempty"`
		PrivilegesRequired    string `json:"privileges_required,omitempty"`
		Score                 int    `json:"score,omitempty"`
		UserInteraction       string `json:"user_interaction,omitempty"`
		Version               int    `json:"version,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
	} `json:"relationships,omitempty"`
}
