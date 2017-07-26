package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if isHost {
		//go run main.go -listen <ip>
		connectionIP := os.Args[2]
		fmt.Println("is host")
		runHost(connectionIP)
	} else {
		//go run main.go <ip>
		connectionIP := os.Args[1]
		fmt.Println("is guest")
		runGuest(connectionIP)
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)

	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: ", message)
}

func runGuest(ip string) {

}