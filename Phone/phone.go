package Phone

import "fmt"

type Camera struct {
	Megapixels int
	Zoom bool
}

type Screen struct {
	ScreenSize int
}
type Phone struct {
	Make string
	Model string
	Camera
	Screen
}

func (p Phone) PrintSummary(){
	fmt.Println("Make: ", p.Make)
	fmt.Println("Model: ", p.Model)
	fmt.Println("MP: ", p.Camera.Megapixels)
	fmt.Println("Zoom: ", p.Camera.Zoom)
	fmt.Println("Screen: ", p.Screen.ScreenSize)
}