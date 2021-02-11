package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("Himanshu Ranjan")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("End of the program.")
}
