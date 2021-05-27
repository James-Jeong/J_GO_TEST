package UdpManager

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient(ip string) {
	fmt.Printf("Starting client... [%s]\n\n", ip)

	conn, err := net.Dial("udp", ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			os.Exit(1)
		}
	}

}
