package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// FetchTargets returns a list of targets belonging to the organizations of the user.
func (a *API) FetchTargets(ctx context.Context, queryOptions *api.TargetQuery,
	pageOptions *api.PageOptions) (targets api.TargetsResponse, offset int, err error) {
	var response api.TargetsResponse

	path := fmt.Sprintf(
		`/targets?fields[organization]=%s&fields[target]=%s`+
			`&page[limit]=%d&page[offset]=%d&filter[organization_id]=%s&include=%s`+
			`&filter[target_group_id]=%s`,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Target,
		pageOptions.GetLimit(),
		pageOptions.GetOffset(),
		queryOptions.Filter.OrganizationID,
		queryOptions.Include,
		queryOptions.Filter.TargetID,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.TargetsResponse{}, 0, err
	}

	if response.Links.Next != "" {
		offset = pageOptions.GetOffset() + api.DefaultPageOffsetIncrement
	}

	return response, offset, nil
}
