package main

import (
	"fmt"
	"time"
)

// <-

// data <- chanVar

// chanVar <- data to write

func SayHello(data chan bool){
	fmt.Println(" say hello go routine")
	time.Sleep(5*time.Second)

	data <- true

}


func main(){

	done := make(chan bool)

	go SayHello(done)

	x := <- done
	fmt.Println(x)
	fmt.Println("Main go routine.")


}
