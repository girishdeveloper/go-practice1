package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func SimpleTCPServer() {
	var TCPPort = ""
	TCPPort = strings.Join(os.Args[1:], "")
	fmt.Printf("Input port is %s\n", ":"+TCPPort)
	listener, err := net.Listen("tcp", ":"+TCPPort)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("Listening to port :%s\n", TCPPort)
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("Error accepting the connection", err1)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())
	if _, err := io.Copy(conn, conn); err != nil {
		fmt.Println("Error echoing data:", err)
	}
}

func main() {
	// Usage: go run TCPserver.go 8080
	SimpleTCPServer()
}
