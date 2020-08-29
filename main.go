package main

import (
	"fmt"
	"log"
	"os"
)

type member struct {
	name     string
	age      int
	location string
}

func main() {
	fmt.Println("Works")
	f := openFile()

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	member := prepareStruct("danny")
	writeFile(f, member)
}

func openFile() *os.File {
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func writeFile(f *os.File, member *member) {

	for i := 0; i <= 5; i++ {
		_, err := f.Write([]byte(member.name + "\n" + member.location + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func prepareStruct(name string) *member {
	member := member{name: name}
	member.location = "Freiburg"
	member.age = 2321
	return &member
}
