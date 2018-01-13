package api

import (
	"errors"

	"github.com/learning-microservice/account/domain"
)

type Service interface {
	Health() string
	Login(name, pass string) (domain.Account, error)
	Register(username, email, password string) (domain.Account, error)
}

type service struct {
	repo domain.AccountRegistory
}

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

func NewService(repo domain.AccountRegistory) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Health() string {
	return "OK"
}

func (s *service) Login(name, password string) (domain.Account, error) {
	a, err := s.repo.GetAccountByName(name)
	if err != nil {
		return domain.Account{}, ErrUnauthorized
	}
	if a.Password != calculatePassHash(password) {
		return domain.Account{}, ErrUnauthorized
	}
	return domain.Account{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
	}, nil
}

func (s *service) Register(username, email, password string) (domain.Account, error) {
	a := domain.Account{}
	a.Username = username
	a.Email = email
	a.Password = calculatePassHash(password)
	err := s.repo.CreateAccount(&a)
	return a, err
}
