package service

import (
	"github.com/atjulia/payment-gateway/gateway-api/internal/domain"
	"github.com/atjulia/payment-gateway/gateway-api/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountResponse, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}
	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByID(ID string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}
