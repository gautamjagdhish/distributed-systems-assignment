package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("\n CLIENT \n")
	// socket
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	for {
		fmt.Print("CLIENT(you) : ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text) //send to server

		p := make([]byte, 2048)
		_, err = bufio.NewReader(conn).Read(p) //recieve from server
		if err == nil {
			fmt.Printf("Server : %s", p)
		} else {
			fmt.Printf("Error : %v\n", err)
		}
	}
	conn.Close()
}
