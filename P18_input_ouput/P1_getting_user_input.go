package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	var age int

	fmt.Printf("Enter your age")

	fmt.Scan(&age)

	fmt.Printf("%T\n", age)
	fmt.Println("my age is ", age)


	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("My name is ", name)


}
