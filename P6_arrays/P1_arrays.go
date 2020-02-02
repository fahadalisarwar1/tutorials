package main

import "fmt"

func main(){
	var arr1 [3]float64
	arr1[0] = 23
	arr1[1] =345
	arr1[2] = 34
	arr1[2] = arr1[0]

	fmt.Println(arr1)
}
