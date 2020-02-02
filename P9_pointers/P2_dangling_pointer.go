package main

import "fmt"

func main(){
	var ptr *string
	if ptr == nil{
		fmt.Println("Its a dangling pointer")
	}
	var x string
	x = "fahad"
	ptr = &x
	//fmt.Printf("type of ptr is : %T\n and value is %v", ptr, ptr)
	//fmt.Printf("type of x is : %T\n and value is %v\n", x, x)
	*ptr = "sarwar"
	fmt.Println(x)


}
