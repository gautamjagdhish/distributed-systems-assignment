package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, msg string) {
	_, err := conn.WriteToUDP([]byte(msg), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	fmt.Println("\n SERVER \n")

	//socket
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	for {
		p := make([]byte, 2048)
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Client : %s", p)
		if err != nil {
			fmt.Printf("Error : %v", err)
			continue
		}

		fmt.Printf("SERVER(you) : ")
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		go sendResponse(ser, remoteaddr, msg)
	}
}
