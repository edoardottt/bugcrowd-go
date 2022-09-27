# bugcrowd-go
Golang Bugcrowd API client

Simple Golang client you can use in order to interact with Bugcrowd API. See the [API documentation here](https://docs.bugcrowd.com/api/2021-10-28/) 

Inspired by [liamg/hackerone](https://github.com/liamg/hackerone).

Usage
-------
```Go
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
```


License 📝
-------

This repository is under [MIT License](https://github.com/edoardottt/bugcrowd-go/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.