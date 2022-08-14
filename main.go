// pbcopy < output.txt | ./main ms username
// pbcopy < `AccountManager ms password`

package main

import (
	"fmt"
	"os"

	Json "encoding/json"
	Acc "main/packages/account"
	File "main/packages/myfile"
	user "os/user"
)

var DirPath = ""
var PathToAccounts = ""
var PathToOutput = ""

func main() {
	var args = os.Args
	var user, _ = user.Current()

	DirPath = user.HomeDir + "/.AcountManager"

	setOutputPath()
	setAccountsPath()
	createApplicationDir()
	createAccountsFile()

	if len(args) <= 1 {
		fmt.Println("Try Commands: ")
		printCommands()
		return
	}

	var allArgsAfter []string = os.Args[1:]
	var firstCmd = allArgsAfter[0]

	if firstCmd == "list" {
		listAccounts(allArgsAfter)
		return
	}

	if len(allArgsAfter) < 2 {
		fmt.Println("More Arguments required")
		return
	}

	if firstCmd == "add" {
		addAccount(allArgsAfter)
		return
	}

	if firstCmd == "update" {
		updateAccount(allArgsAfter)
		return
	}

	if firstCmd == "delete" {
		deleteAccount(allArgsAfter)
		return
	}

	handleAccountRequest(allArgsAfter)
}

func printCommands() {
	var output = `
	list    - List All Accounts
	add     - Add an Account
	update  - Update Account by Account Type
	`

	fmt.Println(output)
}

func listAccounts(args []string) {
	var accounts = getAllAccounts()
	var output string = ""

	for _, account := range accounts {
		output += "Type: " + account.Type + " | Username: " + account.Username + "\n"
	}

	fmt.Println(output)
}

func deleteAccount(args []string) {
	var accType = args[1]
	var allAccounts = getAllAccounts()
	var newAccounts []Acc.Account

	for _, account := range allAccounts {
		if account.Type != accType {
			newAccounts = append(newAccounts, account)
		}
	}

	var res, _ = Json.Marshal(newAccounts)
	os.WriteFile(PathToAccounts, res, 0644)
}

// Full Command
// ./Main add microsoft {username} {password}
func addAccount(args []string) {
	var accType = args[1]
	var username = args[2]
	var password = args[3]

	var allAccounts = getAllAccounts()
	var newAccount = Acc.Account{
		Type:     accType,
		Username: username,
		Password: password,
	}

	allAccounts = append(allAccounts, newAccount)

	var res, _ = Json.Marshal(allAccounts)
	os.WriteFile(PathToAccounts, res, 0644)
}

// Full Command
// ./Main update microsoft {username} {password}
func updateAccount(args []string) {
	var accType = args[1]
	var username = args[2]
	var password = args[3]

	var newAccount = Acc.Account{
		Type:     accType,
		Username: username,
		Password: password,
	}

	var allAccounts = getAllAccounts()

	for index, account := range allAccounts {
		if account.Type == accType {
			allAccounts[index] = newAccount
		}
	}

	var res, _ = Json.Marshal(allAccounts)
	os.WriteFile(PathToAccounts, res, 0644)
}

func handleAccountRequest(args []string) {
	var accountType = args[0]
	var key = args[1]
	var allAccounts = getAllAccounts()
	var account, accountErr = Acc.FindAccount(allAccounts, accountType)

	if accountErr != nil {
		fmt.Println(accountErr)
		return
	}

	if key == "username" {
		File.WriteToFile(PathToOutput, account.Username)
		fmt.Println(PathToOutput)
	}

	if key == "password" {
		File.WriteToFile(PathToOutput, account.Password)
		fmt.Println(PathToOutput)
	}
}

func setOutputPath() {
	PathToOutput = DirPath + "/output.txt"
}

func setAccountsPath() {
	PathToAccounts = DirPath + "/accounts.json"
}

func createApplicationDir() {
	if !File.PathExists(DirPath) {
		File.CreatePath(DirPath)
	}
}

func createAccountsFile() {
	if !File.PathExists(PathToAccounts) {
		var value = make([]Acc.Account, 0)
		var res, _ = Json.Marshal(value)
		os.WriteFile(PathToAccounts, res, 0644)
	}
}

func getAllAccounts() []Acc.Account {
	var file = File.ReadFile(PathToAccounts)

	if file != nil {
		return parseJson(file)
	}

	return nil
}

func parseJson(b []byte) []Acc.Account {
	var accounts []Acc.Account

	Json.Unmarshal(b, &accounts)

	return accounts
}
