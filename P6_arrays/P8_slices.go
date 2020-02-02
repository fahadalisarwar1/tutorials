package main

import "fmt"

func main(){
	arr1 := [...]float64{100, 200, 300, 400, 500}
	slc := arr1[1:4]
	fmt.Println(slc)
	slc[0] = 1
	slc[1] = 2
	slc[2] = 3

	fmt.Println("slice: ",slc)
	fmt.Println("orignal: ", arr1)
	for _, value := range slc{
		fmt.Println(value)
	}
}