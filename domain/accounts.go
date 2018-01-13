package domain

type Account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-",omitempty"`
}
