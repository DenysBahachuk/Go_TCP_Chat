package main

type database interface {
	Open() error
	Save() error
	GetAccountsInfo() map[string]string
	AddAccount(name string, password string)
	ChangeName(oldName string, newName string, password string)
}
