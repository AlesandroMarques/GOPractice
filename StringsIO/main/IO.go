package main

import (
	"fmt"
	"strings"
	"sort"
	"os"
	"log"
	"io/ioutil"


)

func main()  {

	// String stuff
	sampString:= "Hello World"
	// returns true if phrase exists in string
	fmt.Println(strings.Contains(sampString, "lo"))

	// Returns the index for the match
	fmt.Println(strings.Index(sampString, "lo"))

	// Returns the number of matches for the string
	fmt.Println(strings.Count(sampString, "l"))

	// Replaces the first letter with the second as many times as you define
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	// Return a list seperating with defined seperator
	csvString :="1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	listOfStrings := strings.Split(csvString, ",")

	for _, val := range listOfStrings{
		fmt.Println(val)
	}

	//sort stuff
	listOfLetters := []string{"a","c","b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters ", listOfLetters)

	// Returns a string using the values passed in separated with separator
	listOfNums := strings.Join([]string{"3","2","1"}, ",");
	fmt.Println(listOfNums)



	io()


}


func io(){
	file , err := os.Create("samp.txt")

	if err !=nil {

		log.Fatal(err)
	}
	file.WriteString("hellloo")
	file.Close()

	stream , err := ioutil.ReadFile("samp.txt")
	if err !=nil {
		log.Fatal(err)

	}

	readString := string(stream)
	fmt.Println(readString)


}