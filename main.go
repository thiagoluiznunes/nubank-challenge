package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"nubank/domain"
	"nubank/domain/entity"
	"nubank/domain/service"
	"os"
)

func main() {

	var account entity.Account
	var lineMap map[string]interface{}

	accountService := service.AccountService{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), &lineMap)
		if err != nil {
			fmt.Println(err)
		}
		if val, ok := lineMap[domain.Account]; ok {
			_ = accountService.Initialize(&account, val)
		} else if val, ok = lineMap[domain.Transaction]; ok {
			_ = accountService.ProcessTransaction(&account, val)
		}
		err = accountService.PrintAccount(&account)
		if err != nil {
			fmt.Println(err)
		}
		account.Violations = make([]string, 0)
		lineMap = nil
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
