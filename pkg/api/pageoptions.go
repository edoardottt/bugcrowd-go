package api

// PageOptions allows the consumer to request a specific page number and/or size.
type PageOptions struct {
	Limit  int
	Offset int
}

// GetLimit returns the limit with corrections applied.
func (p *PageOptions) GetLimit() int {
	if p == nil || p.Limit <= 0 || p.Limit > 100 {
		return 100
	}
	return p.Limit
}

// GetOffset returns the offset with corrections applied.
func (p *PageOptions) GetOffset() int {
	if p == nil || p.Offset <= 0 {
		return 0
	}
	return p.Offset
}
