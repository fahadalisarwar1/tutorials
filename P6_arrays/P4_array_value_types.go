package main

import "fmt"

func main(){
	arr1 := [...]int64{1,2,3,5,}
	arr2 := arr1

	arr2[2] = 100
	fmt.Println(arr1)
	fmt.Println(arr2)

}
