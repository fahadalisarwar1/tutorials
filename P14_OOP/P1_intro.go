package main

import (
	"fmt"
	"github.com/fahadalisarwar1/tutorials/student"
)

func main(){

	s2 := student.New("abc", "xyz", 25, 350, 500)
	fmt.Println(s2.GetPercentageMarks())
}
