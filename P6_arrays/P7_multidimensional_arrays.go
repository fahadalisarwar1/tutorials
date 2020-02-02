package main

import "fmt"

func main(){
	var twoDimArray [2][3]float64
	twoDimArray[1][2] = 45
	twoDimArray[0][1] = 100
	//fmt.Println(twoDimArray)

	for i, content := range twoDimArray{
		//fmt.Println(i, " ", content)
		for j, value := range content{
			fmt.Printf("i: %v j:%v value: %v\n", i, j, value)
		}
		fmt.Println()
	}

}
