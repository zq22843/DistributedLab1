package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
	fmt.Printf("Enter Text: ")
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	for {
		fmt.Printf("Enter Text: ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		read(conn)
	}

}
