package main

import (

   "log"
   "net/rpc"
//    "os"
   "time"
//    "fmt"
)

type Request struct {
	Data int
	TransNo string
}

type Response struct {
   Data string
   TransNo string
}

var client *rpc.Client

func runGetBalance(dt string) *Response {
	args := new(Request)
	args.TransNo = dt
	response := new(Response)
	divCall := client.Go("Listener.GetBalance", args, response, nil)
	if divCall == nil {
		log.Fatal(divCall)
	}
	return response
}

func runDepositAmount(addmoney int, dt string) *Response {
	args := new(Request)
	args.TransNo = dt
	args.Data = addmoney
	response := new(Response)
	divCall := client.Go("Listener.DepositAmount", args, response, nil)
	if divCall == nil {
		log.Fatal(divCall)
	}
	return response
}

func getTimeNow() string {
	return time.Now().Format("02-01-2006 15:04:05 Monday")
}

func main() {
	client, _ = rpc.DialHTTP("tcp", "127.0.0.1:1234")

	// SET 1
	// request 1
	dt1 := getTimeNow()	
	r1 := runGetBalance(dt1)
	time.Sleep(1 * time.Second)
	// request 2
	dt2 := getTimeNow()		
	r2 := runGetBalance(dt2)
	time.Sleep(1 * time.Second)
	log.Printf("Response1: Transaction Number - %v : %v", r1.TransNo, r1.Data)
	log.Printf("Response2: Transaction Number - %v : %v", r2.TransNo, r2.Data)

	// SET 2
	// request 3
	dt3 := getTimeNow()	
	r3 := runGetBalance(dt3)
	time.Sleep(1 * time.Second)
	// request 4
	dt4 := dt3	
	r4 := runGetBalance(dt4)
	time.Sleep(1 * time.Second)
	log.Printf("Response3: Transaction Number - %v : %v", r3.TransNo, r3.Data)
	log.Printf("Response4: Transaction Number - %v : %v", r4.TransNo, r4.Data)
	

	// SET 3
	// request 5
	dt5 := getTimeNow()	
	addmoney5 := 14675
	r5 := runDepositAmount(addmoney5, dt5)
	time.Sleep(1 * time.Second)
	// request 6
	dt6 := getTimeNow()		
	r6 := runGetBalance(dt6)
	time.Sleep(1 * time.Second)
	log.Printf("Response5: Transaction Number - %v : %v", r5.TransNo, r5.Data)
	log.Printf("Response6: Transaction Number - %v : %v", r6.TransNo, r6.Data)


	// SET 4
	// request 7
	dt7 := getTimeNow()
	addmoney7 := 2500
	r7 := runDepositAmount(addmoney7, dt7)
	time.Sleep(1 * time.Second)
	log.Printf("Response7: Transaction Number - %v : %v", r7.TransNo, r7.Data)
	// request 8
	dt8 := "21-09-2020 20:18:02 Monday"
	addmoney8 := 2500
	r8 := runDepositAmount(addmoney8, dt8)
	time.Sleep(1 * time.Second)
	// request 9
	dt9 := getTimeNow()	
	r9 := runGetBalance(dt9)
	time.Sleep(1 * time.Second)
	log.Printf("Response8: Transaction Number - %v : %v", r8.TransNo, r8.Data)
	log.Printf("Response9: Transaction Number - %v : %v", r9.TransNo, r9.Data)
}
