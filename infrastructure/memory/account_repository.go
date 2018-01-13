package memory

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/learning-microservice/account/domain"
)

var (
	ErrNotFound = errors.New("Notfound")
)

type accountRepositoryOnMemory struct {
	mtx      sync.RWMutex
	accounts map[string]*domain.Account
}

func NewAccountRepository() domain.AccountRegistory {
	return &accountRepositoryOnMemory{
		accounts: make(map[string]*domain.Account),
	}
}

func (r *accountRepositoryOnMemory) GetAccountByName(name string) (*domain.Account, error) {
	for _, a := range r.accounts {
		if a.Username == name {
			return a, nil
		}
	}
	return nil, ErrNotFound
}
func (r *accountRepositoryOnMemory) CreateAccount(a *domain.Account) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	a.ID = fmt.Sprint(time.Now().UnixNano())
	r.accounts[a.ID] = a
	return nil
}
