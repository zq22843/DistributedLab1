package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}

	//TODO In a continuous loop, read a message from the server and display it.
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter Text: ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, _ := net.Dial("tcp", *addrPtr)
	go read(conn)
	write(conn)
}
