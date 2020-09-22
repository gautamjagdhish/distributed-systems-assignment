package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func inputline() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func send(conn net.Conn) {
	fmt.Println("Enter the message to UDP server")
	text := inputline()
	fmt.Fprintf(conn, text) // stream text to conn
	fmt.Println()
}

func receive(conn net.Conn) {
	p := make([]byte, 2048) //empty array init
	//NewReader() returns a new Reader whose buffer has the default size
	//Read() reads data into p. It returns the number of bytes read into p
	_, err := bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("SERVER RESPONSE : ")
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Error2 : %v\n", err)
	}
}

func main() {
	//connect
	addr := "10.250.1.101:8090"
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Println("Error1 : %v\n", err)
	} else {
		fmt.Println("Connected to UDP Server", addr)
	}

	send(conn)
	receive(conn)

	conn.Close()
}

// My name is GautamJagdhish, My roll no is 170010031, What is the sum of numbers in my roll no?
