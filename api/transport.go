package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s Service, logger log.Logger) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	healthHandler := httptransport.NewServer(
		makeHealthEndpoint(s),
		decodeHealthRequest,
		encodeResponse,
		opts...,
	)
	loginHandler := httptransport.NewServer(
		makeLoginEndpoint(s),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	)
	registerHandler := httptransport.NewServer(
		makeRegisterEndpoint(s),
		decodeRegisterRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/account/health", healthHandler).Methods("GET")
	r.Handle("/account/login", loginHandler).Methods("POST")
	r.Handle("/account/register", registerHandler).Methods("POST")

	return r
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// All of our response objects are JSON serializable, so we just do that.
	w.Header().Set("Content-Type", "application/hal+json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	switch err {
	case ErrUnauthorized:
		code = http.StatusUnauthorized
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/hal+json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err.Error(),
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}
