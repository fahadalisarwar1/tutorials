package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
// usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
// Deferred calls are executed even when the function panics


func main(){

	fmt.Printf("This is a main subroutine")

	defer fmt.Println("This is the deffered statement")

	panic("Stop")

	fmt.Println("This is the statement after ")
}


