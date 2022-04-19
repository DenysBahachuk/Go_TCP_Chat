package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type User struct {
	conn       net.Conn
	name       string
	password   string
	room       *Room
	commandsCh chan Command
}

func (u *User) handleConn() {
	for {
		message, err := bufio.NewReader(u.conn).ReadString('\n')
		if err != nil {
			log.Println("Problem with reading the connection", err)
			return
		}
		message = strings.Trim(message, "\r\n")
		args := strings.Split(message, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/register":
			u.commandsCh <- Command{
				id:   CMD_REGISTER,
				user: u,
				args: args,
			}
		case "/login":
			u.commandsCh <- Command{
				id:   CMD_LOGIN,
				user: u,
				args: args,
			}
		case "/list_rooms":
			u.commandsCh <- Command{
				id:   CMD_LIST_ROOMS,
				user: u,
				args: args,
			}
		case "/join":
			u.commandsCh <- Command{
				id:   CMD_JOIN,
				user: u,
				args: args,
			}
		case "/write_to":
			u.commandsCh <- Command{
				id:   CMD_WRITE_TO,
				user: u,
				args: args,
			}
		case "/list":
			u.commandsCh <- Command{
				id:   CMD_LIST,
				user: u,
				args: args,
			}
		case "/list_accounts":
			u.commandsCh <- Command{
				id:   CMD_LIST_ACCOUNTS,
				user: u,
				args: args,
			}
		case "/change_name":
			u.commandsCh <- Command{
				id:   CMD_CHANGE_NAME,
				user: u,
				args: args,
			}
		case "/msg":
			u.commandsCh <- Command{
				id:   CMD_MSG,
				user: u,
				args: args,
			}
		case "/quit_room":
			u.commandsCh <- Command{
				id:   CMD_QUIT_ROOM,
				user: u,
				args: args,
			}
		case "/quit":
			u.commandsCh <- Command{
				id:   CMD_QUIT,
				user: u,
				args: args,
			}
		default:
			u.err(fmt.Errorf("Unknown command %s", cmd))
		}
	}
}

func (u *User) err(err error) {
	u.conn.Write([]byte("Error: " + err.Error() + "\n"))
}

func (u *User) msg(msg string) {
	u.conn.Write([]byte("> " + msg + "\n"))
}
