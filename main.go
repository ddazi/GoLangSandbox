package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Works")
	f := openFile()

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	writeFile(f)
}

func openFile() *os.File {
	f, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func writeFile(f *os.File) {
	for i := 0; i <= 5; i++ {
		_, err := f.Write([]byte("Writing Data:\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
