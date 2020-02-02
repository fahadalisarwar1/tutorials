package main

import "fmt"

func main(){
	sum := 1
	for sum < 100{
		//sum = sum + sum
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("Total: ", sum)

}
