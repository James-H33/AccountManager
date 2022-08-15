package account

import (
	"fmt"
	"os"

	Json "encoding/json"
	File "main/packages/myfile"
)

type Account struct {
	Type     string
	Username string
	Password string
}

type AccountController struct {
	DataPath string
}

func (acc AccountController) Find(accounts []Account, accountType string) (Account, error) {
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

func (acc AccountController) GetAll() []Account {
	var file = File.ReadFile(acc.DataPath)

	if file != nil {
		return parseAccountsJson(file)
	}

	return nil
}

func (acc AccountController) UpdateAccounts(accounts []Account) {
	var res, _ = Json.Marshal(accounts)
	os.WriteFile(acc.DataPath, res, 0644)
}

// Full Command
// ./Main update microsoft {username} {password}
func (acc AccountController) UpdateAccount(args []string) {
	var accType = args[1]
	var username = args[2]
	var password = args[3]

	var newAccount = Account{
		Type:     accType,
		Username: username,
		Password: password,
	}

	var allAccounts = acc.GetAll()

	for index, account := range allAccounts {
		if account.Type == accType {
			allAccounts[index] = newAccount
		}
	}

	acc.UpdateAccounts(allAccounts)
}

func (acc AccountController) ListAccounts(args []string) {
	var accounts = acc.GetAll()
	var output string = ""

	for _, account := range accounts {
		output += "Type: " + account.Type + " | Username: " + account.Username + "\n"
	}

	fmt.Println(output)
}

func (acc AccountController) DeleteAccount(args []string) {
	var accType = args[1]
	var accounts = acc.GetAll()
	var newAccounts []Account

	for _, account := range accounts {
		if account.Type != accType {
			newAccounts = append(newAccounts, account)
		}
	}

	acc.UpdateAccounts(newAccounts)
}

// Full Command
// ./Main add microsoft {username} {password}
func (acc AccountController) AddAccount(args []string) {
	var accType = args[1]
	var username = args[2]
	var password = args[3]

	var accounts = acc.GetAll()
	var newAccount = Account{
		Type:     accType,
		Username: username,
		Password: password,
	}

	accounts = append(accounts, newAccount)
	acc.UpdateAccounts(accounts)
}

func parseAccountsJson(b []byte) []Account {
	var accounts []Account

	Json.Unmarshal(b, &accounts)

	return accounts
}
