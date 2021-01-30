package entity

import (
	"encoding/json"
	"log"
	"time"
)

// Transaction represents a entity
type Transaction struct {
	Merchant string    `json:"merchant,omitempty"`
	Amount   int64     `json:"amount,omitempty"`
	Time     time.Time `json:"time,omitempty"`
}

// ParseLinetoTransction parse line-transaction
func (tx *Transaction) ParseLinetoTransction(line interface{}) (err error) {

	bytes, err := json.Marshal(line)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, tx)
	if err != nil {
		log.Println(err)
	}

	return nil
}
