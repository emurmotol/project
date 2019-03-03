// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeGetByUsernameHandler(m, endpoints, options["GetByUsername"])
	return m
}
