package main

import "fmt"

func main(){
 	num1 := 100
 	fmt.Println("num1 before passing to a function: ", num1)
 	modifyVariable(&num1)
 	//simplefunc(num1)
 	fmt.Println("num1 after passing to the function: ", num1)




 }

 func modifyVariable(ptr *int){
 	*ptr = 200
 }

 func simplefunc(num int){
 	num = 200
 }