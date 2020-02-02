package main

import (
	"fmt"
	"math/rand"
	"time"
)

func subroutine(n int){
	for j := 0; j<= 10; j++{
		fmt.Println(n, ": ", j)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond* amt)
	}

}

func main(){
	for i:= 1; i <= 10; i++{
		go subroutine(i)
		fmt.Printf(" subroutine %v \n", i)
	}

	var input string

	fmt.Scanln(&input)


}