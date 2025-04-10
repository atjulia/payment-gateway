package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("Account not found")
	ErrDuplicatedAPIKey   = errors.New("API Key already exist")
	ErrInvoiceNotFound    = errors.New("Invoice not found")
	ErrUnauthorizedAccess = errors.New("Unauthorized access")
)
