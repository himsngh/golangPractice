package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say : %v\n", ln)
	}
	defer conn.Close()
	// we never reach here // but it turns out we are able to reach at this

	fmt.Println("Done handling ..... ")
}
