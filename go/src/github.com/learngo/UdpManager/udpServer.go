package UdpManager

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func StartServer(ip string) {
	fmt.Printf("Starting server... [%s]\n\n", ip)

	udpAddr := GetUdpAddr(ip)
	udpConn := Listen(udpAddr)

	for {
		display(udpConn)
	}
}

func display(udpConn *net.UDPConn) {
	netData, err := bufio.NewReader(udpConn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.TrimSpace(string(netData)) == "STOP" {
		fmt.Println("Exiting TCP server!")
		os.Exit(1)
	}

	fmt.Print("-> ", string(netData))
	t := time.Now()
	myTime := t.Format(time.RFC3339) + "\n"
	fmt.Fprintf(udpConn, myTime+"\n")
}
