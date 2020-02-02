package main

import "fmt"

func main(){
	var x1  float64 = 33
	var x2 float64 = 55

	s, _ := Calculator(x1, x2)
	fmt.Println("Sum is : ", s)
	fmt.Println("diff is : " )

}

func Calculator(num1, num2 float64)(sum, diff float64){
	sum = num1 + num2
	diff = num1 - num2
	return sum, diff
}