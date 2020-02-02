package main

import "fmt"

func main(){
	arr := [3]int{10, 20, 30,}
	fmt.Println("Original Array: ", arr)
	passingSlices(arr[:])
	fmt.Println("After passing: ", arr)

}


func passingSlices(slc []int){
	slc[1] = 99
}