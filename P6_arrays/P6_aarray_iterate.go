package main

import "fmt"

func main(){
	var arr = [5]float64{11, 22, 33, 44, 55,}
	for _, content := range arr{
		fmt.Println(content)

	}

	str1 := [...]string{"My", "name", "is", "Fahad"}
	for i, value := range str1{
		fmt.Println(i, " ", value)
	}


}
