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
	outputArray()
	outputSliceTestingReference()
	creatingAndLoopMap()
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

func outputArray() {
	newArray := [5]string{"test1", "test2", "test3", "test4", "test5"}
	fmt.Println("Content Array Complete", newArray)

	for i, v := range newArray {
		fmt.Println("Content Array Line ", i, v)

	}
}

func outputSliceTestingReference() {
	newSlice := []string{"testSlice1", "testSlice2", "testSlice3", "testSlice4", "testSlice5"}
	fmt.Println("Content Slice Complete", newSlice)

	for i, v := range newSlice {
		fmt.Println("Content Slice Line ", i, v)
	}
	newSliceReferenz := newSlice

	newSliceReferenz[0] = "testSliceChanged"
	fmt.Println("Slice Reference:", newSliceReferenz)
	fmt.Println("Slice Origin", newSlice)

}

func creatingAndLoopMap() {
	m := make(map[string]string)

	m["test1"] = "blub1"
	m["test2"] = "blub2"
	m["test3"] = "blub3"
	fmt.Println("Looping through Map:")

	for i, v := range m {
		fmt.Println("Key:", i)
		fmt.Println("Value:", v)
	}
}
