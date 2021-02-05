package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo() // this will not be executed in the case of log.Fatal().
	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("Error fmt : \t", err)
		//log.Fatal("Error log Fatal \t", err)  // after log fatal nothing will run and it will exit the program with exit status 1.
		log.Println("Error log Println \t ", err)
		panic(err)
	}
	fmt.Println("Reached end of line.") // this will not execute in both the condition with panic or log.Fatal.

	defer foo() // this defer won't run in the panicking call.
}

func foo() {
	fmt.Println("\n Foo bar is deferred from main")
}
