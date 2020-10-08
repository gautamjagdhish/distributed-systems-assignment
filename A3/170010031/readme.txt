Assignment 3 – Clock Synchronization using Berkeley Algorithm

Objective : To understand about clock synchronization in a distributed setting by implementing Berkeley Algorithm.

Input files :
master-client.go : to run master and slave depending on argument
slaves.txt : this file contains details about the slave nodes. each slave is in ip:port format, space seperated

Output :
4 shiviz compatibe govec log files are generated - 1 master node log and 3 slave node logs
master-Log.txt, s6666-Log.txt, s7777-Log.txt, s8888-Log.txt
All the four logs are combined together in shiviz.txt. It's then uploaded to the shiviz website to visualise the logs.

Screenshots : 
terminals.png contains the screenshot of 4 terminals showing the update of time in all the nodes
shiviz.png contains the screenshot of the visualisation generated using the log files

How it works :
First, the master sends it's time to all the slaves.
Now, each of the slaves recieve the master's time, calculates the difference between their time and master time and sends the difference back to master
The master recieves the differences from all the slaves and calculates the average difference of all the clocks.
Now, master adjusts its own time and tells each slave how to adjust their time.
Finally, the slaves update their clocks and the time is synchronised in the system.
The above operation happens every 5s. So the clock is synchronised every 5s

Command line arguments/Flags :
To ease the use of command line arguments, I used the flags library in which each of the argument can be specified in "-name=value" format
Specify the flags in following format,
-m or -m=True for master node
-s or -s=True for client node
-a=ip_address:port for master and client
-t=intial_time for master and client
-sf=slavesfile for master
-lf=logfile for master and client

How to run :
Check terminals.png for reference
Master : go run master-client.go -m -a=127.0.0.1:1111 -t=15 -sf=slaves.txt -lf=master
Slaves(according to slaves.txt) :
go run master-client.go -s -a=127.0.0.1:6666 -t=35 -lf=s6666 
go run master-client.go -s -a=127.0.0.1:7777 -t=-10 -lf=s7777 
go run master-client.go -s -a=127.0.0.1:8888 -t=20 -lf=s8888


