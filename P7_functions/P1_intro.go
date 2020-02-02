package main

import "fmt"

func main(){
	var num1 float64 = 44
	var num2 float64 = 55
	av := findAverage(num1, num2)
	fmt.Println(av)

}


func findAverage(num1, num2 float64)(average float64){
	average = (num1 + num2)/2
	return
}
