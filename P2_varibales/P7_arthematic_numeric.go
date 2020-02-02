package main

import "fmt"

import "go"
func main(){
	var num1 float64
	var num2 float64
	num1 = 39
	num2 = 45
	fmt.Printf("%v + %v = %v\n", num1, num2, (num1+num2))
	fmt.Printf("%v - %v = %v\n", num1, num2, (num1-num2))
	fmt.Printf("%v * %v = %v\n", num1, num2, (num1*num2))
	fmt.Printf("%v / %v = %v\n", num1, num2, (num1/num2))
}
