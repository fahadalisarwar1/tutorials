package main

import "fmt"

// The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual
// function call is done.

func printValueofA(a int){
	fmt.Println("The value of a in defered statement is ", a)
}

func main(){

	a := 100

	defer printValueofA(a)

	a = 200
	fmt.Println("value of a before deferred function call is ", a)


}
