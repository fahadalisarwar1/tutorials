package main

import "fmt"

func main(){
	a := 1 + 2i
	fmt.Println(a)

	b := complex(1,99)
	fmt.Println(b)

	c := a + b
	fmt.Println(c)

	d := a * b

	fmt.Println(d)

}
