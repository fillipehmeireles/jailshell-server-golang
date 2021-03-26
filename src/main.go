package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"tcp-server/src/server"
	"tcp-server/src/syscore"
)

func init() {
	syscore.SetupCore()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please, set a port number")
		fmt.Println("E.g.: ./luspew-jailshell 4000")
		return
	}

	port := ":" + arguments[1]

	listener, erro := net.Listen("tcp4", port)
	if erro != nil {
		log.Fatal(erro)
	}
	defer listener.Close()

	fmt.Printf("Starting server on port: %s \n", port)

	for {
		c, erro := listener.Accept()
		if erro != nil {
			log.Fatal(erro)
		}

		go server.HandleConnection(c)
	}
}
