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

var BasePath = ""
var PathToAccounts = ""
var PathToOutput = ""
var AccControl Acc.AccountController

func main() {
	var args []string = os.Args[1:]

	setBasePath()
	setOutputPath()
	setAccountsPath()
	createApplicationDir()
	createAccountsFile()

	AccControl = Acc.AccountController{
		DataPath: PathToAccounts,
	}

	if len(args) < 1 {
		fmt.Println("Try Commands: ")
		listCommands()
		return
	}

	var firstCmd = args[0]

	if firstCmd == "list" {
		AccControl.ListAccounts(args)
		return
	}

	if len(args) < 2 {
		fmt.Println("More Arguments required")
		return
	}

	if firstCmd == "add" {
		AccControl.AddAccount(args)
		return
	}

	if firstCmd == "update" {
		AccControl.UpdateAccount(args)
		return
	}

	if firstCmd == "delete" {
		AccControl.DeleteAccount(args)
		return
	}

	handleGetAccountInfoRequest(args)
}

func listCommands() {
	var output = `
	list    - List All Accounts
	add     - Add an Account
	update  - Update Account by Account Type
	delete  - Delete Account by Account Type
	`

	fmt.Println(output)
}

func handleGetAccountInfoRequest(args []string) {
	var accountType = args[0]
	var key = args[1]
	var accounts = AccControl.GetAll()
	var account, accountErr = AccControl.Find(accounts, accountType)

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
	PathToOutput = BasePath + "/output.txt"
}

func setBasePath() {
	var user, _ = user.Current()
	BasePath = user.HomeDir + "/.AcountManager"
}

func setAccountsPath() {
	PathToAccounts = BasePath + "/accounts.json"
}

func createApplicationDir() {
	if !File.PathExists(BasePath) {
		File.CreatePath(BasePath)
	}
}

func createAccountsFile() {
	if !File.PathExists(PathToAccounts) {
		var value = make([]Acc.Account, 0)
		var res, _ = Json.Marshal(value)
		os.WriteFile(PathToAccounts, res, 0644)
	}
}
