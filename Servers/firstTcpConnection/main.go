package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(conn, "\nCalling from Tcp connection ........ \n")
		fmt.Fprintln(conn, " Hello World form tcp server.... Hurray !")
		fmt.Fprintf(conn, "%v", "You are doing great just keep going \n")
		conn.Close()
	}
}
