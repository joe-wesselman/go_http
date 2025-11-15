package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"log"
	"os"
)

func main(){
	fmt.Println("hello world")
	protocol := "tcp"
	port := ":8000"

	ln, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for {
		msg, err := bufio.NewReader(conn).ReadString("\n")
		if err != nil{
			log.Fatal(err)
		}
		fmt.Print("Message received:", string(msg))
	}
}