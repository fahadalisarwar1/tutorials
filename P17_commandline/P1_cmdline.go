package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
 	args  := os.Args
 	num1 := args[1]
 	num2 := args[2]
 	fmt.Println(num1 + num2)

 	num3, _ := strconv.Atoi(num1)
	num4, _ := strconv.Atoi(num2)
	fmt.Println(num3 + num4)


 	fmt.Println(args[1])

}
