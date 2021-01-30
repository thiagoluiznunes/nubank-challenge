package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"nubank/domain"
	"nubank/domain/entity"
	"time"
)

// AccountService represents a service entity
type AccountService struct{}

// Initialize initialize a Account
func (s *AccountService) Initialize(account *entity.Account, line interface{}) (err error) {

	if !account.Active {
		account.ParseLineToAccount(line)
		account.Violations = make([]string, 0)
		account.Active = true
	} else {
		account.Violations = append(account.Violations, domain.AccountAlreadyInitialized)
		return errors.New(domain.AccountAlreadyInitialized)
	}

	return nil
}

// ProcessTransaction procces one transaction
func (s *AccountService) ProcessTransaction(account *entity.Account, line interface{}) (err error) {

	var transaction entity.Transaction
	if account.Active && account.ActiveCard {
		err = transaction.ParseLinetoTransction(line)
		if err != nil {
			return err
		}
		err = s.validateDuplicateTransactions(account, &transaction)
		if err != nil {
			return err
		}
		err = s.validateMultipleTransactions(account, &transaction)
		if err != nil {
			return err
		}
		err = s.validateLimitUpdate(account, &transaction)
		if err != nil {
			return err
		}
	} else if !account.Active {
		account.Violations = append(account.Violations, domain.AccountNotInitialized)
		return errors.New(domain.AccountNotInitialized)
	} else {
		account.Violations = append(account.Violations, domain.CardNotActive)
		return errors.New(domain.CardNotActive)
	}

	if len(account.Violations) > 0 {
		return errors.New(domain.ContainViolations)
	}

	account.Transactions = append(account.Transactions, transaction)
	account.AvailableLimit -= transaction.Amount

	return nil
}

func (s *AccountService) validateMultipleTransactions(account *entity.Account, transaction *entity.Transaction) (err error) {

	if len(account.Transactions) > 2 {
		var txsWithinTwoMinutes []entity.Transaction
		for _, value := range account.Transactions {
			if s.validateTimeBetween(transaction.Time, value.Time) {
				txsWithinTwoMinutes = append(txsWithinTwoMinutes, value)
			}
			if len(txsWithinTwoMinutes) > 2 {
				account.Violations = append(account.Violations, domain.HighFrequencySmallInterval)
				break
			}
		}
	}

	return nil
}

func (s *AccountService) validateDuplicateTransactions(account *entity.Account, transaction *entity.Transaction) (err error) {

	for _, value := range account.Transactions {
		if value.Merchant == transaction.Merchant &&
			value.Amount == transaction.Amount &&
			s.validateTimeBetween(transaction.Time, value.Time) {
			account.Violations = append(account.Violations, domain.DoubleTransaction)
			break
		}
	}

	return nil
}

func (s *AccountService) validateTimeBetween(txTime time.Time, check time.Time) bool {

	startTime := txTime.Add(time.Duration(-2) * time.Minute)
	endTime := txTime.Add(time.Second)

	return check.After(startTime) && check.Before(endTime)
}

func (s *AccountService) validateLimitUpdate(account *entity.Account, transaction *entity.Transaction) (err error) {

	if account.AvailableLimit-transaction.Amount < 0 {
		account.Violations = append(account.Violations, domain.InsufficientLimit)
	}

	return nil
}

// PrintAccount prints
func (s *AccountService) PrintAccount(account *entity.Account) (err error) {

	bytes, err := json.Marshal(account)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))

	return nil
}
