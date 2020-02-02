package main

// Go Workspace : /home/fahad/go/src

import (
	"fmt"
	"github.com/fahadalisarwar1/tutorials/geometry"

)



func main(){
	var l, w float64
	l = 34
	w = 45
	rectArea := geometry.Area(l, w)
	fmt.Println(rectArea)

}
