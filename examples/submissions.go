package main

import (
	"context"
	"fmt"

	bugcrowd "github.com/edoardottt/bugcrowd-go"
	"github.com/edoardottt/bugcrowd-go/pkg/api"
)

func main() {
	b := bugcrowd.New("username", "token")

	queryOptions := api.SubmissionQuery{}
	pageOptions := api.PageOptions{}
	submissions, _, err := b.Services.FetchSubmissions(context.TODO(), &queryOptions, &pageOptions)

	if err != nil {
		fmt.Println(err)
	}

	for _, submission := range submissions.Data {
		fmt.Println(submission.ID)
	}
}
