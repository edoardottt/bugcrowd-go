package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// FetchOrganization returns a single organization by UUID.
func (a *API) FetchOrganization(ctx context.Context,
	queryOptions *api.OrganizationQuery) (organization api.OrganizationResponse,
	err error) {
	var response api.OrganizationResponse

	path := fmt.Sprintf(
		`/organizations/%s?fields[organization]=%s&fields[program]=%s`+
			`&fields[target]=%s&include=%s`,
		queryOptions.ID,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Program,
		queryOptions.Fields.Target,
		queryOptions.Include,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.OrganizationResponse{}, err
	}

	return response, nil
}

// FetchOrganizations returns a list of organizations belonging to the user.
func (a *API) FetchOrganizations(ctx context.Context, queryOptions *api.OrganizationQuery,
	pageOptions *api.PageOptions) (organizations api.OrganizationsResponse, offset int, err error) {
	var response api.OrganizationsResponse

	path := fmt.Sprintf(
		`/organizations?fields[organization]=%s&fields[program]=%s`+
			`&fields[target]=%s&page[limit]=%d&page[offset]=%d&include=%s`+
			`&sort=%s`,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Program,
		queryOptions.Fields.Target,
		pageOptions.GetLimit(),
		pageOptions.GetOffset(),
		queryOptions.Include,
		queryOptions.Sort,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.OrganizationsResponse{}, 0, err
	}

	if response.Links.Next != "" {
		offset = pageOptions.GetOffset() + api.DefaultPageOffsetIncrement
	}

	return response, offset, nil
}
