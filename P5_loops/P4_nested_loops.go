package main

import "fmt"

func main(){
	num1 := 5
	for i := 0; i < num1; i++{
		for j := 0; j < i; j++{
			fmt.Printf("i = %d j = %d ", i, j)
		}
		fmt.Println()
	}
}