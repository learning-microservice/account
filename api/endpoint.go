package api

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/learning-microservice/account/domain"
)

type healthRespose struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

func makeHealthEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		status := s.Health()
		return healthRespose{
			Service: "account",
			Status:  status,
			Time:    time.Now().String(),
		}, nil
	}
}

type loginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type loginRespose struct {
	Account domain.Account `json:"account"`
}

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		a, err := s.Login(req.Username, req.Password)
		return loginRespose{Account: a}, err
	}
}

type registerRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
type registerRespose struct {
	Account domain.Account `json:"account"`
}

func makeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		a, err := s.Register(req.Username, req.Email, req.Password)
		return registerRespose{Account: a}, err
	}
}
