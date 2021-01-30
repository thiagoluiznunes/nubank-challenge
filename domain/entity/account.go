package entity

import (
	"encoding/json"
	"log"
)

// Account represents a entity
type Account struct {
	Active         bool          `json:"-"`
	ActiveCard     bool          `json:"active-card"`
	AvailableLimit int64         `json:"available-limit"`
	Violations     []string      `json:"violations"`
	Transactions   []Transaction `json:"-"`
}

// ParseLineToAccount parse line-account
func (account *Account) ParseLineToAccount(line interface{}) (err error) {

	bytes, err := json.Marshal(line)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(bytes, account)
	if err != nil {
		log.Println(err)
	}

	return nil
}
