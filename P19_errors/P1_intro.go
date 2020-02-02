package main

import "fmt"

type MyError struct {

}

func (myErr *MyError) Error()string{
	return "Something wrong happened with the program"
}

func main(){

	myErr := &MyError{}
	fmt.Println(myErr)

}
