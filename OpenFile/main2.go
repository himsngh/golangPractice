package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Main 2")
	file, err := os.Open("/Users/himanshsingh/Desktop/golangPractice/FileOs/names.txt")

	// data := make([]byte, 100)

	// count, _ := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("read %d bytes: %q \n", count, data[:count])

	bs, _ := ioutil.ReadAll(file)
	fmt.Print(string(bs))
	file.Close()
	// checking the goto statement

	fmt.Println("\nGoto statemetn")
	count := 0

	// har baar initialize ho rha hai ye dhyan rakhna hai
LOOP:
	for count < 10 {
		count++
		if count == 5 {
			goto LOOP
		}
		fmt.Println(count)
	}
	// LOOP:
	// 	for i := 0; i < 1; i++ {
	// 		if count == 10 {
	// 			break
	// 		}
	// 		for j := 0; j < 1; j++ {
	// 			fmt.Printf("i == %d ,j == %d and count == %d ", i, j, count)
	// 			if i == j {
	// 				count++
	// 				goto LOOP
	// 			}
	// 			fmt.Printf("%d != %d", i, j)
	// 		}

	// 	}
}
