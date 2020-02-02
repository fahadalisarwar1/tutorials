package main

import "fmt"

func main(){
	chanVar := make(chan int)

	fmt.Printf("%T ", chanVar)
	fmt.Println()

}