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
)

func main() {
	b := bugcrowd.New("your-username-here", "your-password-here")
	submissions, _, _ := b.Services.FetchSubmissions(context.TODO(), nil)
	for _, submission := range submissions.Data {
		fmt.Println(submission.Id)
	}
}
```