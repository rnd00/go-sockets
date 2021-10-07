package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"go-sockets/s/helpers"
)

const (
	cHost = "localhost"
	cPort = "27314"
	cType = "tcp"
)

func main() {
	fmt.Println("Starting " + cType + " server on " + cHost + ":" + cPort)
	l, e := net.Listen(cType, cHost+":"+cPort)
	if e != nil {
		fmt.Println("Error listening:", e.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println("Error connecting:", e.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	buffer, e := bufio.NewReader(c).ReadBytes('\n')

	if e != nil {
		fmt.Println("Client left.")
		c.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	helpers.ParseFnCall(buffer[:len(buffer)-1])

	c.Write(buffer)

	handleConnection(c)
}
