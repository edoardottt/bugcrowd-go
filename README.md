# bugcrowd-go
Simple Golang client to interact with Bugcrowd API. See the [API documentation here](https://docs.bugcrowd.com/api/2021-10-28/).  

<a href="https://edoardoottavianelli.it">
	<img src="https://github.com/edoardottt/bugcrowd-go/actions/workflows/go.yml/badge.svg" alt="workflows" />
</a>

Usage 🚀
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

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/bugcrowd-go/issues) / [pull request](https://github.com/edoardottt/bugcrowd-go/pulls).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run
```bash
golangci-lint run
```
If there aren't errors, go ahead :)

Inspired by [liamg/hackerone](https://github.com/liamg/hackerone).

License 📝
-------

This repository is under [MIT License](https://github.com/edoardottt/bugcrowd-go/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
