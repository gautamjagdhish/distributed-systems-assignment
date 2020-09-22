Assignment 1 – UDP Programming using Golang

Objective : To get familiar with the Go language and implement a client program that communicates with the server using UDP.

How to run : go run udpclient.go

How it works : First, the program connects to the server from main(). Next, it calls send() which inputs a line from STDIN and sends it to the server. Now, receive() is called which waits for response from the server and prints it.