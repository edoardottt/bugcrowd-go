package services

import (
	"github.com/edoardottt/bugcrowd-go/pkg/client"
)

type API struct {
	client *client.Client
}

func New(c *client.Client) *API {
	return &API{
		client: c,
	}
}
