package api

type Links struct {
	Self     string `json:"self,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}
