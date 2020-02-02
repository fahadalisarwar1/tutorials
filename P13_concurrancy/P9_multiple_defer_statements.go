package main

import "fmt"

func main(){
	fmt.Println("Hello")
	defer fmt.Println("outside loop statement")
	for i:=1; i<=5; i++{
		defer fmt.Println(i)
	}
	fmt.Println("function returned")

}