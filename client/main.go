package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	cHost = "localhost"
	cPort = "27314"
	cType = "tcp"
)

func main() {
	fmt.Println("Connecting to", cType, "server", cHost+":"+cPort)
	c, e := net.Dial(cType, cHost+":"+cPort)
	if e != nil {
		fmt.Println("Error connecting:", e.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")

		input, _ := reader.ReadString('\n')

		c.Write([]byte(input))

		msg, _ := bufio.NewReader(c).ReadString('\n')

		log.Print("Server relay: " + msg)
	}
}
