package integration

import (
	"encoding/json"
	"nubank/domain"
	"nubank/domain/entity"
	"nubank/domain/service"
	"testing"
)

var line string
var lineMap map[string]interface{}
var mockAccount entity.Account
var accountService service.AccountService

// AccountNotInitialized try make a trasaction without account
func AccountNotInitialized(t *testing.T) {

	line = `{"transaction": {"merchant": "Burger King 65", "amount": 10, "time":"2019-02-13T10:45:05.000Z"}`
	json.Unmarshal([]byte(line), &lineMap)

	// should contain account not initialized violation
	err := accountService.ProcessTransaction(&mockAccount, lineMap[domain.Transaction])
	if err.Error() == domain.AccountNotInitialized {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
	} else {
		t.Error(err)
		return
	}
}

// InitializeAccount init account
func InitializeAccount(t *testing.T) {

	mockAccount = entity.Account{}
	accountService = service.AccountService{}

	line = `{"account": {"active-card": true, "available-limit": 100}}`
	json.Unmarshal([]byte(line), &lineMap)

	// should not contain violations
	err := accountService.Initialize(&mockAccount, lineMap[domain.Account])
	if err == nil && len(mockAccount.Violations) <= 0 {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
	} else {
		t.Error(err)
		return
	}
}

// AccountAlreadyInitialized try to init account
func AccountAlreadyInitialized(t *testing.T) {

	line = `{"account": {"active-card": true, "available-limit": 100}}`
	json.Unmarshal([]byte(line), &lineMap)

	// should contain account already initilized violation
	err := accountService.Initialize(&mockAccount, lineMap)
	if len(mockAccount.Violations) > 0 && mockAccount.Violations[0] == domain.AccountAlreadyInitialized {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
	} else {
		t.Error(err)
		return
	}
}
