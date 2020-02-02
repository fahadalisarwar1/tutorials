package main

import "fmt"

func main(){
	num1 := 150
	num2 := 101
	if num1 < num2{
		fmt.Println(num1, " < ", num2)
	}else if num2 == 100{
		fmt.Println(num1, " = ", num2)
	}else{
		fmt.Println(num2)
	}


}
