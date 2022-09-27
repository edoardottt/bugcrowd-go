package bugcrowd

import (
	"github.com/edoardottt/bugcrowd-go/pkg/client"
	"github.com/edoardottt/bugcrowd-go/pkg/services"
)

type API struct {
	Services *services.API
}

func New(username, token string) *API {
	c := client.New(username, token)

	return &API{
		Services: services.New(c),
	}
}
