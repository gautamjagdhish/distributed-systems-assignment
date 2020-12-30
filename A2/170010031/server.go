package main
import (
	"fmt"
	// "log"
	"net/http"
	"net/rpc"
	"os"
	"io/ioutil"
	"time"
	// "bufio"
	"strconv"
	"strings"
)

type Listener int

type Request struct {
	Data int
	TransNo string
}

type Response struct {
   Data string
   TransNo string
}

var TransIds []string
var index int
var file []byte

func (l *Listener) GetBalance(args *Request, response *Response) error {
	
	tno := args.TransNo

	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	for i := range TransIds {
		if TransIds[i] == tno {	
			response.Data = "The transaction already processed"
			response.TransNo = tno
			return nil
		} 
	}	

	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file
	TransIds = append(TransIds, tno)
	f, _ := os.OpenFile("Trans_Processed.txt", os.O_APPEND|os.O_WRONLY, 0644)
	_, _ = fmt.Fprintln(f, tno)
	f.Close()

	//Implement the logic to read the Balance from Balance.txt
	file, _ := ioutil.ReadFile("Balance.txt")
	bal := strings.Split(string(file), "\n")[0]

	//Implement the logic to add Balance and Transaction ID to the response
	response.Data = bal
	response.TransNo = tno

	return nil
}

func (l *Listener) DepositAmount(args *Request, response *Response) error {

	tno := args.TransNo
	newbal := args.Data

	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	for i := range TransIds {
		if TransIds[i] == tno {	
			response.Data = "The transaction already processed"
			response.TransNo = tno
			return nil
		} 
	}	

	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file
	TransIds = append(TransIds, tno)
	f1, _ := os.OpenFile("Trans_Processed.txt", os.O_APPEND|os.O_WRONLY, 0644)
	_, _ = fmt.Fprintln(f1, tno)
	f1.Close()

	//Implement the logic to read the balance and add the amount to into balance and write calculated balance back to Balance.txt
	data, _ := ioutil.ReadFile("Balance.txt")
	bal := strings.Split(string(data), "\n")[0]
	curbal, _ := strconv.Atoi(bal)
	balbal := strconv.Itoa(curbal + newbal)
    _ = ioutil.WriteFile("Balance.txt", []byte(balbal), 0644)

	// add delay
	time.Sleep(3 * time.Second)
	
	//Implement the logic add appropriate message and transaction number to response	
	response.Data = "Your Amount is deposited into the account successfully"
	response.TransNo = tno

	return nil
}

func main() {

	//Read the transaction processed from Trans_Processed.txt into TransIds
	file, _ = ioutil.ReadFile("Trans_Processed.txt")
	TransIds = strings.Split(string(file), "\n")

	//Sample code to start the server, read the port number from command-line
	listener := new(Listener)
	rpc.Register(listener)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

