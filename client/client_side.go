package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:4545")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Connection closed!")
		done <- struct{}{}
	}()
	io.Copy(conn, os.Stdin)
	conn.Close()
	<-done
}
