package main

import "fmt"

func main(){
	var summ float64

	summ = summation(1,2,3,4,56,7)
	summ = summation(100,200,300)
	fmt.Println(summ)
}

func summation(inputParam ...float64)(sum float64){
	sum = 0
	for _, entry := range inputParam{
		sum += entry
	}
	return sum
}
