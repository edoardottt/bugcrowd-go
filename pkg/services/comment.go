package services

import (
	"context"

	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

// CreateComment creates a comment on a submission.
func (a *API) CreateComment(ctx context.Context, dataOptions *api.CommentData) (comment api.Comment, err error) {
	var response api.Comment

	path := `/comments`
	if err := a.client.Post(ctx, path, &response, dataOptions); err != nil {
		return api.Comment{}, err
	}

	return response, nil
}
