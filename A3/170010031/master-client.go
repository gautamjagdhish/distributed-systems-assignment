package main

import (
    "flag"
	"fmt"
    // "os"
	"io/ioutil"
	"strings"
	t "time"
	"net"
    "strconv"
    // "reflect"
    
    
	gv "github.com/DistributedClocks/GoVector/govec"
)

var(
    logger *gv.GoLog
    config = gv.GetDefaultConfig()
    options = gv.GetDefaultLogOptions()
    // maddress string
) 

/*
func ttt(address string, time int, slavesfile, logfile string) {
    fmt.Println(address, time, slavesfile, logfile)
}
*/

func printError(err error) {
    if err != nil {
        fmt.Println(err)
        return
    }
}

func master(address string, time float64, slavesfile, logfile string) {


    // read slave addresses
    data, _ := ioutil.ReadFile(slavesfile)
    slaves := strings.Split(string(data), " ")

    logger = gv.InitGoVector("Master", logfile, config)

    var timediffs = make(map[string]float64)

    for {
        fmt.Println("Master time :", time)
        for _, slave := range slaves {
            // connect to slave address
            // if(slave == "127.0.0.1:6666") {
            //     fmt.Println("same")
            // } else {
            //     fmt.Println("not same")
            // }

            udpaddr, err := net.ResolveUDPAddr("udp", slave)
            printError(err)
            conn, err := net.DialUDP("udp", nil, udpaddr)
            printError(err)
            defer conn.Close()

            // send time to slave
            msg := []byte(fmt.Sprintf("%f", time))
            encmsg := logger.PrepareSend("sending master time", msg, options)
            _, err = conn.Write(encmsg)
            printError(err)
            fmt.Println("sent master time to", slave)
        
            // recieve diff from slave
            encmsg2 := make([]byte, 1024)
            var  msg2 string
            _, err = conn.Read(encmsg2)
            logger.UnpackReceive("time diff received", encmsg2, &msg2, options)
            fmt.Println("recieved time diff", msg2, "from", slave)
            timediffs[slave], _ = strconv.ParseFloat(msg2, 64)
            // printError(err)
            
        }

        // cal average of all slaves and master
        var sum float64
        sum = 0
        for _, v := range timediffs {
            sum += v
            // fmt.Println(v, sum)
        }
        avg := sum/float64(len(timediffs)+1)

        fmt.Println("average diff is", avg)
        fmt.Println()
        // update server time
        time = time + avg  

        for address, diff := range timediffs {
            udpaddr, err := net.ResolveUDPAddr("udp", address)
            printError(err)
            conn, err := net.DialUDP("udp", nil, udpaddr)
            printError(err)
            defer conn.Close()

            msg := []byte(fmt.Sprintf("%f", avg-diff))
            encmsg := logger.PrepareSend("sending correction", msg, options)
            _, err = conn.Write(encmsg)
        }
        t.Sleep(5*t.Second)
     }
}


func slave(address string, time float64, logfile string) {
    // listen on slave's address 
    udpaddr, err := net.ResolveUDPAddr("udp", address)
    printError(err)
    conn, err := net.ListenUDP("udp", udpaddr)
    printError(err)
    defer conn.Close()

    logger = gv.InitGoVector("slave@"+address, logfile, config)
    encmsg := make([]byte, 1024)
    var msg string
    
    // wait to receive msg
    for {
        // master sends time to slave
        _, addr, err := conn.ReadFromUDP(encmsg)
        printError(err)
        logger.UnpackReceive("time received from master", encmsg, &msg, options)
        mtime, err := strconv.ParseFloat(msg, 64)
        printError(err)
        diff := time - mtime
        fmt.Println("master time :", mtime, "slave time :", time)

        // slave replies the time diff
        msg2 := []byte(fmt.Sprintf("%f", diff))
        encmsg = logger.PrepareSend("time diff sent", msg2, options)
        _, err = conn.WriteToUDP(encmsg, addr)
        printError(err)
        fmt.Println("time diff", diff, "sent to master")

        // master replies with adjustment
        _, addr, err = conn.ReadFromUDP(encmsg)
        printError(err)
        logger.UnpackReceive("adjustment received from master", encmsg, &msg, options)
        adj, err := strconv.ParseFloat(msg, 64)
        printError(err)
        fmt.Println("adjustment received", adj)
        time = time + adj
        fmt.Println("time updated to", time)
        fmt.Println()
        
    
        t.Sleep(5*t.Second)
    }
}



func main() {
    var (
        isMaster, isSlave *bool
        address, slavesfile, logfile *string
        time *float64
    )

    isMaster = flag.Bool("m", false, "is master ?")
    isSlave = flag.Bool("s", false, "is slave ?")
    address = flag.String("a", "127.0.0.1:1234", "ip:port")
    time = flag.Float64("t", 0, "time")
    slavesfile = flag.String("sf", "slavesfile.txt", "slavesfile")
    logfile = flag.String("lf", "logfile.txt", "logfile")
    flag.Parse()

    if(*isMaster == true) {
        master(*address, *time, *slavesfile, *logfile)
    } else if(*isSlave == true) {
        slave(*address, *time, *logfile)
    }
}

