package main

import (
	"fmt"
)

func sayHello(){
	fmt.Println("This is a sub routine!")

}

func main(){
	fmt.Println("This is main go routine ")

	go sayHello()

	//time.Sleep(3*time.Second)
	fmt.Println("this is a statement after the go routine ")
}
