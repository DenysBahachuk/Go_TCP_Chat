package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	db "github.com/DenysBahachuk/databaseExample"
)

type Server struct {
	port       string
	rooms      map[string]*Room
	commandsCh chan Command
	accauntsDB database
	countUsers int
}

func newServer() *Server {

	return &Server{
		port:       ":4545",
		rooms:      make(map[string]*Room),
		commandsCh: make(chan Command),
		accauntsDB: db.NewDataBase(),
		countUsers: 0,
	}
}

func (s *Server) run() {
	s.accauntsDB.Open()
	for cmd := range s.commandsCh {
		switch cmd.id {
		case CMD_REGISTER:
			s.register(cmd.user, cmd.args)
		case CMD_LOGIN:
			s.login(cmd.user, cmd.args)
		case CMD_LIST_ROOMS:
			s.listRooms(cmd.user)
		case CMD_JOIN:
			s.join(cmd.user, cmd.args)
		case CMD_WRITE_TO:
			s.writeTo(cmd.user, cmd.args)
		case CMD_LIST:
			s.list(cmd.user)
		case CMD_LIST_ACCOUNTS:
			s.listAccounts(cmd.user)
		case CMD_CHANGE_NAME:
			s.changeName(cmd.user, cmd.args)
		case CMD_MSG:
			s.message(cmd.user, cmd.args)
		case CMD_QUIT_ROOM:
			s.quitRoom(cmd.user)
		case CMD_QUIT:
			s.quit(cmd.user)
		}
	}
}

func (s *Server) newUser(conn net.Conn) *User {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())
	s.countUsers++
	return &User{
		conn:       conn,
		name:       "Anonymous_" + strconv.Itoa(s.countUsers),
		commandsCh: s.commandsCh,
	}
}

func (s *Server) login(u *User, args []string) {

	if len(args) < 3 {
		u.msg("Name and password are required.")
		return
	}
	u.name = args[1]
	u.password = args[2]

	pass, ok := s.accauntsDB.GetAccountsInfo()[u.name]
	if !ok || pass != u.password {
		u.msg("Wrong name or password. Try again")
		return
	}
	u.msg(fmt.Sprintf("Welcome to chat, %s!", u.name))
}

func (s *Server) register(u *User, args []string) {

	if len(args) < 3 {
		u.msg("Name and password are required.")
		return
	}

	name := args[1]
	password := args[2]

	if _, ok := s.accauntsDB.GetAccountsInfo()[name]; ok {
		u.msg("Registration failed. The name or password are already exist.")
		return
	}

	s.accauntsDB.AddAccount(name, password)
	u.msg("Registration success. You may login now.")
}

func (s *Server) listRooms(u *User) {
	if len(s.rooms) == 0 {
		u.msg("There are no available rooms.")
	} else {
		u.msg("Available rooms:")
		for _, r := range s.rooms {
			u.msg(r.name)
		}
	}
}

func (s *Server) join(u *User, args []string) {

	var roomName string
	if len(args) < 2 {
		roomName = "Common"
	} else {
		roomName = args[1]
	}
	s.quitRoom(u)
	r, ok := s.rooms[roomName]
	if !ok {
		r = &Room{
			name:    roomName,
			members: make(map[net.Addr]*User),
		}
		s.rooms[roomName] = r
	}
	r.members[u.conn.RemoteAddr()] = u
	u.room = r
	r.broadcast(u, fmt.Sprintf("%s has joined.", u.name))
	u.msg(u.name + " welcome to " + roomName + " room.")
}

func (s *Server) writeTo(u *User, args []string) {

	receiverName := args[1]

	if len(s.rooms) == 0 {
		u.msg("There are no active users!")
	}
	var receiver *User
	for _, room := range s.rooms {
		for _, member := range room.members {
			if receiverName == member.name {
				receiver = member
			}
		}
	}
	message := strings.Join(args[2:], " ")
	receiver.msg(message)
}

func (s *Server) list(u *User) {
	if len(s.rooms) == 0 {
		u.msg("There are no active users!")
	}
	for _, room := range s.rooms {
		for _, member := range room.members {
			u.msg(member.name)
		}
	}
}

func (s *Server) listAccounts(u *User) {
	for name, _ := range s.accauntsDB.GetAccountsInfo() {
		u.msg(name)
	}
}

func (s *Server) changeName(u *User, args []string) {
	if len(args) < 4 {
		u.msg("Name, password and a newName are required.")
		return
	}

	name := args[1]
	password := args[2]
	newName := args[3]

	pass, ok := s.accauntsDB.GetAccountsInfo()[name]
	if !ok || pass != password {
		u.msg("Wrong name or password. Try again")
		return
	}
	s.accauntsDB.ChangeName(name, newName, password)
	u.name = newName
	u.msg("Name succesfully changed!")
}

func (s *Server) message(u *User, args []string) {
	u.room.broadcast(u, u.name+": "+strings.Join(args[1:], " "))
}

func (s *Server) quitRoom(u *User) {
	if u.room != nil {
		delete(u.room.members, u.conn.RemoteAddr())
		u.room.broadcast(u, fmt.Sprintf("%s has left the room", u.name))
	}
}

func (s *Server) quit(u *User) {
	log.Printf("Client %s has disconnected", u.conn.RemoteAddr().String())
	s.quitRoom(u)
	u.msg("See you later!")
	s.accauntsDB.Save()
	u.conn.Close()
}
