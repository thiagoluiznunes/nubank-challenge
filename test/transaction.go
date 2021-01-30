package integration

import (
	"encoding/json"
	"fmt"
	"nubank/domain"
	"testing"
)

// InsufficientLimit try to execute a transaction without sufficient limit
func InsufficientLimit(t *testing.T) {

	line = `{"transaction": {"merchant": "Burger King 65", "amount": 100000, "time":"2019-02-13T10:45:05.000Z"}}`
	json.Unmarshal([]byte(line), &lineMap)

	// should contain insufficient limit violation
	err := accountService.ProcessTransaction(&mockAccount, lineMap[domain.Transaction])
	if err.Error() == domain.ContainViolations && mockAccount.Violations[0] == domain.InsufficientLimit {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
	} else {
		t.Error(err)
		return
	}
}

// ProcessTransaction try to init account
func ProcessTransaction(t *testing.T) {

	line = `{"transaction": {"merchant": "Burger King 65", "amount": 10, "time":"2019-02-13T10:45:05.000Z"}}`
	json.Unmarshal([]byte(line), &lineMap)

	// should not contains violations
	err := accountService.ProcessTransaction(&mockAccount, lineMap[domain.Transaction])
	if err == nil && len(mockAccount.Violations) <= 0 {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
	} else {
		t.Error(err)
		return
	}
}

// HighFrequencySmallInterval try to init account
func HighFrequencySmallInterval(t *testing.T) {

	for i := 0; i < 3; i++ {
		line = fmt.Sprintf(`{"transaction": {"merchant": "Burger King 65", "amount": %d, "time":"2019-02-13T10:45:05.000Z"}}`, i+1)
		json.Unmarshal([]byte(line), &lineMap)

		// should contain hight frequency small interval violation
		err := accountService.ProcessTransaction(&mockAccount, lineMap[domain.Transaction])
		if len(mockAccount.Violations) > 0 && mockAccount.Violations[0] == domain.HighFrequencySmallInterval {
			_ = accountService.PrintAccount(&mockAccount)
			mockAccount.Violations = make([]string, 0)
			return
		} else if err != nil {
			t.Error(err)
			return
		}
	}
}

// CardNotActive try to execute a transaction without sufficient limit
func CardNotActive(t *testing.T) {

	mockAccount.ActiveCard = false
	line = `{"transaction": {"merchant": "Burger King 65", "amount": 100000, "time":"2019-02-13T10:45:05.000Z"}}`
	json.Unmarshal([]byte(line), &lineMap)

	// error should be card-not-active violation
	err := accountService.ProcessTransaction(&mockAccount, lineMap[domain.Transaction])
	if len(mockAccount.Violations) > 0 && mockAccount.Violations[0] == domain.CardNotActive {
		_ = accountService.PrintAccount(&mockAccount)
		mockAccount.Violations = make([]string, 0)
		mockAccount.ActiveCard = true
	} else {
		t.Error(err)
		return
	}
}
