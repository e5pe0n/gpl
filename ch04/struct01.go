package main

type Employee struct {
	ID            int
	Name, Address string
	string
	DoB       string
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000

	position := &dilbert.Position
	*position = "Senior" + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
}
