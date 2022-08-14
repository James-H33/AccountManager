package account

import "fmt"

type Account struct {
	Type     string
	Username string
	Password string
}

func FindAccount(accounts []Account, accountType string) (Account, error) {
	var output Account

	for _, account := range accounts {
		if account.Type == accountType {
			output = account
		}
	}

	if output.Type == "" {
		return output, fmt.Errorf("Account not found")
	}

	return output, nil
}
