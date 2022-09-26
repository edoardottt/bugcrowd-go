package bugcrowd

import (
	"github.com/edoardottt/bugcrowd-go/pkg/client"
	"github.com/edoardottt/bugcrowd-go/pkg/services"
)

type API struct {
	Services *services.API
}

func New(username, password string) *API {

	c := client.New(username, password)

	return &API{
		Services: services.New(c),
	}
}
