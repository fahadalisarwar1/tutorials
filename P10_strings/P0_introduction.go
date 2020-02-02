package main

import "fmt"

func main(){
	var myStr string

	myStr = "seÑor"
	printBytes(myStr)
	fmt.Println(len(myStr))


}

func printBytes(str string){
	for i := 0; i< len(str); i++{
		fmt.Printf("%c ", str[i])
		fmt.Printf("%x ", str[i]) // %x prints hex values
	}
	fmt.Println()
}
