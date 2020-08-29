package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Works")
	f := openFile()

	err := f.Close()

	if err != nil {
		log.Fatal(err)
	}

}

func openFile() *os.File {
	f, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
