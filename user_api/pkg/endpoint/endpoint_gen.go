// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "github.com/emurmotol/project/user_api/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetByUsernameEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.UserApiService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{GetByUsernameEndpoint: MakeGetByUsernameEndpoint(s)}
	for _, m := range mdw["GetByUsername"] {
		eps.GetByUsernameEndpoint = m(eps.GetByUsernameEndpoint)
	}
	return eps
}
