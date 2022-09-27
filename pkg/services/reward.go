package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// FetchReward returns a single monetary reward by UUID.
func (a *API) FetchReward(ctx context.Context, queryOptions *api.MonetaryRewardQuery) (reward api.MonetaryRewardResponse,
	err error) {
	var response api.MonetaryRewardResponse
	path := fmt.Sprintf(
		`/monetary_rewards/%s?fields[identity]=%s&fields[monetary_reward]=%s`+
			`&fields[payment]=%s&fields[submission]=%s&include=%s`,
		queryOptions.Id,
		queryOptions.Fields.Identity,
		queryOptions.Fields.MonetaryReward,
		queryOptions.Fields.Payment,
		queryOptions.Fields.Submission,
		queryOptions.Include,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.MonetaryRewardResponse{}, err
	}

	return response, nil
}

// UpdateSubmission cancels a monetary reward and provide a reason.
func (a *API) UpdateReward(ctx context.Context, queryOptions *api.MonetaryRewardQuery,
	dataOptions *api.MonetaryRewardData) (reward api.MonetaryRewardResponse, err error) {
	var response api.MonetaryRewardResponse
	path := fmt.Sprintf(
		`/monetary_rewards/%s`,
		queryOptions.Id,
	)
	if err := a.client.Patch(ctx, path, &response, dataOptions); err != nil {
		return api.MonetaryRewardResponse{}, err
	}

	return response, nil
}

// CreateReward creates a new monetary reward for a submission.
func (a *API) CreateReward(ctx context.Context, dataOptions *api.MonetaryRewardData) (reward api.MonetaryRewardResponse, err error) {
	var response api.MonetaryRewardResponse
	path := `/monetary_rewards`
	if err := a.client.Post(ctx, path, &response, dataOptions); err != nil {
		return api.MonetaryRewardResponse{}, err
	}

	return response, nil
}
