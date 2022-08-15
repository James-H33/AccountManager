# Go Account Manager Tool

Create Executable
```
go build -o AccountManager
```

Create Symbolic Link - Mac OS
```
ln -s /path/to/{ExecutableName} /path/to/bin
```


Run and Copy output to Clipboard -> Examples:
```
pbcopy < `AccountManager microsoft username`   
pbcopy < `AccountManager microsoft password`   
```

## Data Stored
Data for this application is stored in `/user/{Username}/.AccountManager`

Accounts Data Structure:
```json
{
    "type":  "account-name",
    "username: "username"
    "password: "password"
}
```
