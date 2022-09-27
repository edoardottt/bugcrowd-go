package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// FetchSubmission returns a single submission by UUID.
func (a *API) FetchSubmission(ctx context.Context, queryOptions *api.SubmissionQuery) (submission api.SubmissionResponse,
	err error) {
	var response api.SubmissionResponse
	path := fmt.Sprintf(
		`/submissions/%s?fields[activity]=%s&fields[claim_ticket]=%s&`+
			`fields[comment]=%s&fields[cvss_vector]=%s&`+
			`fields[external_issue]=%s&fields[file_attachment]=%s&`+
			`fields[identity]=%s&fields[monetary_reward]=%s&`+
			`fields[organization]=%s&fields[payment]=%s&`+
			`fields[program]=%s&fields[program_brief]=%s&`+
			`fields[submission]=%s&fields[target]=%s&`+
			`fields[target_group]=%s&include=%s`,
		queryOptions.Id,
		queryOptions.Fields.Activity,
		queryOptions.Fields.ClaimTicket,
		queryOptions.Fields.Comment,
		queryOptions.Fields.CvssVector,
		queryOptions.Fields.ExternalIssue,
		queryOptions.Fields.FileAttachment,
		queryOptions.Fields.Identity,
		queryOptions.Fields.MonetaryReward,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Payment,
		queryOptions.Fields.Program,
		queryOptions.Fields.ProgramBrief,
		queryOptions.Fields.Submission,
		queryOptions.Fields.Target,
		queryOptions.Fields.TargetGroup,
		queryOptions.Include,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.SubmissionResponse{}, err
	}

	return response, nil
}

// UpdateSubmission updates a single submission by UUID.
func (a *API) UpdateSubmission(ctx context.Context, queryOptions *api.SubmissionQuery,
	dataOptions *api.SubmissionData) (submission api.SubmissionResponse, err error) {
	var response api.SubmissionResponse
	path := fmt.Sprintf(
		`/submissions/%s`,
		queryOptions.Id,
	)
	if err := a.client.Patch(ctx, path, &response, dataOptions); err != nil {
		return api.SubmissionResponse{}, err
	}

	return response, nil
}

// FetchSubmissions returns a filtered list of submissions based on tokenized search and sort parameters.
func (a *API) FetchSubmissions(ctx context.Context, queryOptions *api.SubmissionQuery,
	pageOptions *api.PageOptions) (submissions api.SubmissionResponse, offset int, err error) {
	var response api.SubmissionResponse
	path := fmt.Sprintf(
		`/submissions?fields[activity]=%s&fields[claim_ticket]=%s&fields[comment]=%s`+
			`&fields[cvss_vector]=%s&fields[external_issue]=%s&fields[file_attachment]=%s`+
			`&fields[identity]=%s&fields[monetary_reward]=%s&fields[organization]=%s`+
			`&fields[payment]=%s&fields[program]=%s&fields[program_brief]=%s&fields[submission]=%s`+
			`&fields[target]=%s&fields[target_group]=%s&page[limit]=%d&page[offset]=%d`+
			`&include=%s&filter[assignee]=%s&filter[blocked_by]=%s&filter[duplicate]=%t`+
			`&filter[payments]=%s&filter[points]=%s&filter[program]=%s&filter[researcher]=%s`+
			`&filter[retest]=%s&filter[search]=%s&filter[severity]=%s&filter[source]=%s`+
			`&filter[state]=%s&filter[submitted]=%s&filter[target]=%s&filter[target_type]=%s`+
			`&filter[updated]=%s&filter[vrt]=%s&sort=%s`,
		queryOptions.Fields.Activity,
		queryOptions.Fields.ClaimTicket,
		queryOptions.Fields.Comment,
		queryOptions.Fields.CvssVector,
		queryOptions.Fields.ExternalIssue,
		queryOptions.Fields.FileAttachment,
		queryOptions.Fields.Identity,
		queryOptions.Fields.MonetaryReward,
		queryOptions.Fields.Organization,
		queryOptions.Fields.Payment,
		queryOptions.Fields.Program,
		queryOptions.Fields.ProgramBrief,
		queryOptions.Fields.Submission,
		queryOptions.Fields.Target,
		queryOptions.Fields.TargetGroup,
		pageOptions.GetLimit(),
		pageOptions.GetOffset(),
		queryOptions.Include,
		queryOptions.Filters.Assignee,
		queryOptions.Filters.BlockedBy,
		queryOptions.Filters.Duplicate,
		queryOptions.Filters.Payments,
		queryOptions.Filters.Points,
		queryOptions.Filters.Program,
		queryOptions.Filters.Researcher,
		queryOptions.Filters.Retest,
		queryOptions.Filters.Search,
		queryOptions.Filters.Severity,
		queryOptions.Filters.Source,
		queryOptions.Filters.State,
		queryOptions.Filters.Submitted,
		queryOptions.Filters.Target,
		queryOptions.Filters.TargetType,
		queryOptions.Filters.Updated,
		queryOptions.Filters.Vrt,
		queryOptions.Sort,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.SubmissionResponse{}, 0, err
	}

	if response.Links.Next != "" {
		offset = pageOptions.GetOffset() + 50
	}

	return response, offset, nil
}

// CreateSubmission creates a submission within a program.
func (a *API) CreateSubmission(ctx context.Context, dataOptions *api.SubmissionData) (submission api.SubmissionResponse, err error) {
	var response api.SubmissionResponse
	path := `/submissions`
	if err := a.client.Post(ctx, path, &response, dataOptions); err != nil {
		return api.SubmissionResponse{}, err
	}

	return response, nil
}
