package domain

type AccountRegistory interface {
	GetAccountByName(name string) (*Account, error)
	CreateAccount(account *Account) error
}
