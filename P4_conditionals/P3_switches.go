package main

import "fmt"

func main(){
	var option int
	option = 2
	switch option{
	case 1:
		fmt.Println("go running")
	case 2:
		fmt.Println("go jogging")
	case 3:
		fmt.Println("do push ups")
	default:
		fmt.Println("Take Rest now")
	}

}
