package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	endpoint "github.com/emurmotol/project/user_api/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeGetByUsernameHandler creates the handler logic
func makeGetByUsernameHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/username/{username}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetByUsernameEndpoint, decodeGetByUsernameRequest, encodeGetByUsernameResponse, options...)))
}

// decodeGetByUsernameRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByUsernameRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetByUsernameRequest{}
	return req, nil
}

// encodeGetByUsernameResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByUsernameResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
