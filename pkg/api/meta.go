package api

type Meta struct {
	TotalHits int `json:"total_hits,omitempty"`
	Count     int `json:"count,omitempty"`
}
