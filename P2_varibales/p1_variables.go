package main

import "fmt"

// here global scoped

var gpa int
var studentName string
var isGraduated bool


func main(){
	// custom scoped
	gpa = 3
	studentName ="Fahad"
	isGraduated = true
	fmt.Println(gpa)
	fmt.Println(studentName)
	fmt.Println(isGraduated)
	fmt.Println("GPA: ", gpa, " Name:", studentName, " Graduated: ", isGraduated)


}
