Assignment 2 – At-Most-Once Semantics

Objective : to understand and implement At-Most-Once Semantics scenarios using rpc in golang

How to run :
go run server.go
go run client.go

How it works : 
* Once the server is setup, it listens on the specified port
* Now,  when the client program is executed, it sends the RPC calls as specified to the same port the server is listening on
* RPC calls contain 2 arguments : Request and Response
* Request contains the parameters to be passed on to the function and Response is an empty object to be used to transport the output
* The functions in server.go executes using the data provided in Request and attaches the output to Response
* Now if the RPC call was executed successfully, client accesses the Response object and logs it in the console

Functions in server has been implemented as specified

I have used 3 helper functions in client.go :
> runGetBalance : It takes a datetime string as the argument and returns the pointer to the response object
> runDepositAmount : It takes a datetime string and the amount to be added as the arguments and returns the pointer to the response object
> getTimeNow : It returns the current date and time in "02-01-2006 15:04:05 Monday" format