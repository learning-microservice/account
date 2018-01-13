package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/learning-microservice/account/domain"
)

var (
	s = NewService(&mockAccountRepository{})
)

func TestHealth(t *testing.T) {
	actual := s.Health()
	expected := "OK"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLogin(t *testing.T) {
	actual, _ := s.Login("user1", "hogehoge")
	expected := domain.Account{
		ID:       "123456",
		Username: "user1",
		Email:    "user1@dummy.com",
	}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestLogin_InvalidPassword(t *testing.T) {
	_, actual := s.Login("user1", "NG")
	expected := ErrUnauthorized
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestLogin_AccountNotfound(t *testing.T) {
	_, actual := s.Login("user2", "NG")
	expected := ErrUnauthorized
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRegister(t *testing.T) {
	actual, _ := s.Register("user999", "user1@dummy.com", "hogehogehoge")
	expected := domain.Account{
		ID:       "999999",
		Username: "user999",
		Email:    "user1@dummy.com",
		Password: calculatePassHash("hogehogehoge"),
	}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

type mockAccountRepository struct{}

func (r *mockAccountRepository) GetAccountByName(name string) (*domain.Account, error) {
	if name == "user1" {
		return &domain.Account{
			ID:       "123456",
			Username: name,
			Email:    fmt.Sprintf("%s@dummy.com", name),
			Password: calculatePassHash("hogehoge"),
		}, nil
	} else if name == "user2" {
		return nil, errors.New("not found")
	} else {
		panic("invalid")
	}
}
func (r *mockAccountRepository) CreateAccount(a *domain.Account) error {
	a.ID = "999999"
	return nil
}
