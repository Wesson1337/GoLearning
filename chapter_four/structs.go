package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID, Salary, ManagerID int
	Name, Address, Position string
	DoB time.Time
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000 // Зарплата снижена, пишет мало строк кода
	fmt.Println(dilbert)

	position := &dilbert.Position
	*position = "Senior" + *position // Повышен в должности
	fmt.Println(dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (активный участник команды)"
	fmt.Println(dilbert)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // Уволить
	fmt.Println(dilbert)
}

func EmployeeByID(id int) *Employee { /* ... */ return &dilbert}
