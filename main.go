package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type member struct {
	name     string
	age      int
	location string
}
type address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type userData struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  address `json:"address"`
}

func main() {
	fmt.Println("Works")
	f := openFile()

	defer func() {
		err := f.Close()
		deleteFile(f)
		if err != nil {
			log.Fatal(err)
		}
	}()
	member := prepareStruct("danny")
	writeFile(f, member)
	//	outputArray()
	//	outputSliceTestingReference()
	//	creatingAndLoopMap()
	//go createAndServeHttp()
	//	getHttpLocal()
	getJSON()
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

func deleteFile(f *os.File) {
	err := os.Remove(f.Name())
	if err != nil {
		log.Fatal(err)
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

func createAndServeHttp() {
	http.HandleFunc("/foo", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}

}
func hello(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hi I'm from the Webserver\n")
	if err != nil {
		panic(err)
	}

}

func getHttpLocal() {
	resp, err := http.Get("http://localhost:8090/foo")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Response:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func getJSON() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var c []userData
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}

	for l := range c {
		fmt.Printf("Id = %v, Name = %v", c[l].Id, c[l].Username)
		fmt.Println()
	}
}
