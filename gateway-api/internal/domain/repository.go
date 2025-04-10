package domain

type AccountRepository interface {
	Save(Account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(Invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
