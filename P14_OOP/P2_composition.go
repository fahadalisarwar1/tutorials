package main

import "github.com/fahadalisarwar1/tutorials/Phone"

func main(){
	p1 := Phone.Phone{
		Make:   "Oneplus",
		Model:  "3T",
		Camera: Phone.Camera{16, true},
		Screen: Phone.Screen{256},
	}

	p1.PrintSummary()

}
