package main

import (
	"./lib"
	"flag"
	"fmt"
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
		lib.RunHost(connectionIP)
	} else {
		//go run main.go <ip>
		connectionIP := os.Args[1]
		fmt.Println("is guest")
		lib.RunGuest(connectionIP)
	}
}