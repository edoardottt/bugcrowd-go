package services

import (
	"context"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// CreateClaimTicket creates claim ticket with submission relationship.
func (a *API) CreateClaimTicket(ctx context.Context, dataOptions *api.ClaimTicketData) (claimTicket api.ClaimTicket,
	err error) {
	var response api.ClaimTicket

	path := `/claim_tickets`
	if err := a.client.Post(ctx, path, &response, dataOptions); err != nil {
		return api.ClaimTicket{}, err
	}

	return response, nil
}
