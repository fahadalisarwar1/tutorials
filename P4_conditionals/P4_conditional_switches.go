package main

import "fmt"

func main(){
	num1 := 34
	num2 := 100
	switch { // and
	case num1 == 50 && num2 < 100:
		fmt.Println("Number 1 is 50")
	case num1 < 50:
		fmt.Println("Number is less than 50")
	case num1==50 && num2 >=50:
		fmt.Println("Number1 is 50 and number 2 is more than 50")
	default:
		fmt.Println("This is a defautl case")
	}
}
