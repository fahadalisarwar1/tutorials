package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	fmt.Println("This is a subroutine for printing numbers")

	for i:= 1; i<=5; i++{
		// 1 2 3 4 5 numbers to be pinted

		// 1 2 3 4 5 waiting time
		time.Sleep(1*time.Second)

		fmt.Printf("%v ", i)
	}
	fmt.Println("Print numbers subroutine finished")

}

func printAlphabets(){
	fmt.Println("This is subroutine for printing alphabets")

	for i := 'a'; i<= 'e'; i++{
		 // a b c d e
		// 2 4 6  8 10
		 time.Sleep(2*time.Second)

		 fmt.Printf("%c ", i)

	}
	fmt.Println("Alphabets printing subroutine finished")
}



func main(){
	fmt.Println("Main sub routine ")

	go printNumbers()
	go printAlphabets()

	time.Sleep(13*time.Second)
	fmt.Println("Main go routine finsihed")
}