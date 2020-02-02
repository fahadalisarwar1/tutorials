package main

import "fmt"

var (
	name string
	age int
)

func main(){
	lastName := "Sarwar" // type inference
	marks := 50

	fmt.Println(lastName)
	fmt.Println(marks)
	hobby, phoneNum := "playing hockey", "12345"
	fmt.Println(hobby)
	fmt.Println(phoneNum)

	hobby, phoneNum = "abc", "def"
	fmt.Println(hobby)
	fmt.Println(phoneNum)

}