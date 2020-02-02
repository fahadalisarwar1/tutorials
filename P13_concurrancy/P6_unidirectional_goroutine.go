package main

import "fmt"

func greetUser(chanVar chan <- int){
	fmt.Println("Say Greeting to the user")

	chanVar <- 100

}

func main(){
	fmt.Println("Main routine")


	chanVar := make(chan <- int)

	go greetUser(chanVar)



}