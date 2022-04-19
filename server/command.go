package main

type commandID int

const (
	CMD_REGISTER commandID = iota
	CMD_LOGIN
	CMD_LIST_ROOMS
	CMD_JOIN
	CMD_WRITE_TO
	CMD_LIST
	CMD_LIST_ACCOUNTS
	CMD_CHANGE_NAME
	CMD_MSG
	CMD_QUIT_ROOM
	CMD_QUIT
)

type Command struct {
	id   commandID
	user *User
	args []string
}
