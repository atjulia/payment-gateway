package repository

import (
	"database/sql"
	"time"

	"github.com/atjulia/payment-gateway/gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stm, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return err
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE api_key = ?
	`, apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE id = ?
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`
		SELECT balance FROM accounts WHERE id = ? FOR UPDATE
	`, account.ID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE accounts
		SET balance = ?, updated_at = ?
		WHERE id = ?
	`, currentBalance+account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
