package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func TCPclient() {
	var TCPport = os.Args[1]
	var Host = os.Args[2]
	// connect to TCP server on port 8080
	con, err := net.Dial("tcp", Host+":"+TCPport)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer con.Close()
	fmt.Println("Connected to " + Host + " on port " + TCPport)

	//Send some data to server
	reader := bufio.NewReader(os.Stdin)
	for {
		//data := "Hello, server! this is girish"
		fmt.Print("You $ ")
		data, _ := reader.ReadString('\n')
		//fmt.Print(reflect.TypeOf(data), data)
		// Remove the character \n from the input
		data = strings.Replace(data, "\n", "", -1)
		// Check if the input is "exit"
		if strings.Compare(data, "exit") == 0 {
			fmt.Println("Stopping the client")
			break
		}
		if _, err1 := io.Copy(con, strings.NewReader(data)); err1 != nil {
			fmt.Println("Error sending data:", err1)
			return
		}
		// Read the echoed data back from the server
		buf := make([]byte, len(data))
		if _, err2 := io.ReadFull(con, buf); err != nil {
			fmt.Println("Error reading data:", err2)
			return
		}
		fmt.Println("Received from server:", string(buf))
	}
	os.Exit(0)
}

func main() {
	// USAGE: go run TCPclient.go 8080 localhost
	TCPclient()
}
