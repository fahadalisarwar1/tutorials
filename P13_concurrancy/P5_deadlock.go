package main

import (
	"fmt"
	"time"
)

func sayGreeting(){
	fmt.Println("This is a sub routine")

	time.Sleep(5*time.Second)

}


func main(){
	chanVar := make(chan int)
	go sayGreeting()

	x := <- chanVar

	fmt.Println(x)


}