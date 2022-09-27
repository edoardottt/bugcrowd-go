package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// FetchProgram returns a single program by UUID.
func (a *API) FetchProgram(ctx context.Context, queryOptions *api.ProgramQuery) (program api.ProgramResponse,
	err error) {
	var response api.ProgramResponse

	path := fmt.Sprintf(
		`programs/%s?fields[organization]=%s&fields[program]=%s`+
			`&fields[program_brief]=%s&fields[reward_range]=%s`+
			`&fields[submission]=%s&fields[target]=%s`+
			`&fields[target_group]=%s&include=%s`,
		queryOptions.ID,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Program,
		queryOptions.Fields.ProgramBrief,
		queryOptions.Fields.RewardRange,
		queryOptions.Fields.Submission,
		queryOptions.Fields.Target,
		queryOptions.Fields.TargetGroup,
		queryOptions.Include,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.ProgramResponse{}, err
	}

	return response, nil
}

// FetchPrograms returns a list of programs belonging to the user.
func (a *API) FetchPrograms(ctx context.Context, queryOptions *api.ProgramQuery,
	pageOptions *api.PageOptions) (programs api.ProgramsResponse, offset int, err error) {
	var response api.ProgramsResponse

	path := fmt.Sprintf(
		`/programs?fields[organization]=%s&fields[program]=%s`+
			`&fields[program_brief]=%s&fields[reward_range]=%s`+
			`&fields[submission]=%s&fields[target]=%s`+
			`&fields[target_group]=%s&page[limit]=%d`+
			`&page[offset]=%d&include=%s&filter[organization_id]=%s`+
			`&sort=%s`,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Program,
		queryOptions.Fields.ProgramBrief,
		queryOptions.Fields.RewardRange,
		queryOptions.Fields.Submission,
		queryOptions.Fields.Target,
		queryOptions.Fields.TargetGroup,
		pageOptions.GetLimit(),
		pageOptions.GetOffset(),
		queryOptions.Include,
		queryOptions.Filter.OrganizationID,
		queryOptions.Sort,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.ProgramsResponse{}, 0, err
	}

	if response.Links.Next != "" {
		offset = pageOptions.GetOffset() + api.DefaultPageOffsetIncrement
	}

	return response, offset, nil
}
