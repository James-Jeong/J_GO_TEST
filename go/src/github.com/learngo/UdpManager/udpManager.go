package UdpManager

import (
	"fmt"
	"net"
	"strings"
)

func CheckIp(ipAddr string) bool {
	addr := net.ParseIP(ipAddr[:strings.IndexByte(ipAddr, ':')]) // IP:Port 에서 IP 만 Parsing
	if addr == nil {
		fmt.Println("Invalid address")
		return false
	} else {
		fmt.Println("The address is ", addr.String())
	}
	return true
}

func GetUdpAddr(address string) *net.UDPAddr {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return udpAddr
}

func Listen(udpAddr *net.UDPAddr) *net.UDPConn {
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return udpConn
}
