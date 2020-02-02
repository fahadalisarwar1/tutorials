package student

type Student struct {
	// fields
	Name string
	LastName string
	Age float64
	Marks float64
	TotalMarks float64
}

func New(Name string,  LastName string, Age float64, Marks float64, TotalMarks float64)(s Student){
	s = Student{Name:Name, LastName:LastName, Age:Age, Marks:Marks, TotalMarks:TotalMarks}
	return s
}

//
func (s Student) GetPercentageMarks()float64{
	perc := s.Marks/s.TotalMarks * 100
	return perc
}


