package main

import (
	"log"
	"net"
)

func main() {

	server := newServer()
	go server.run()

	listener, err := net.Listen("tcp", server.port)
	if err != nil {
		log.Fatalf("Failed to connect server %s", err.Error())
	}
	defer listener.Close()

	log.Printf("Started server on %s", server.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection %s", err.Error())
			continue
		}

		user := server.newUser(conn)
		conn.Write([]byte("Welcome to chat! Enter a command: \n"))
		go user.handleConn()

	}

}
