package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Empid  int
	Salary float32
}

type Ceo struct {
	Employee
	Emp   Employee
	IsCeo bool
}

func (e Employee) GetSalary() float32 {
	return e.Salary
}

func (e *Employee) SetSalary(s float32) {
	e.Salary = s
}

func main() {
	e := Employee{Name: "bob", Empid: 123, Salary: 99.99}
	fmt.Println(e)
	fmt.Println(e.GetSalary())
	e.SetSalary(123.123)
	fmt.Println(e.GetSalary())
	c := Ceo{
		Employee: Employee{Name: "man", Empid: 1, Salary: 54.32},
		Emp:      Employee{Name: "lady", Empid: 2, Salary: 9092.23},
		IsCeo:    true}
	fmt.Printf("c: %v\n", c)

	// o, _ := json.MarshalIndent(c, "", "    ")
	// fmt.Println(string(o))
	// fmt.Println(o)

	// s := "{\"Name\": \"man\",\"Empid\": 1,\"Salary\": 54.32,\"Emp\": {\"Name\": \"lady\",\"Empid\": 2,\"Salary\": 9092.23},\"IsCeo\": true}"

	// var decoded Ceo
	// decoder := json.NewDecoder(strings.NewReader(s))
	// decoder.Decode(&decoded)
	// fmt.Println(decoded)

	// fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", c.Emp)
	// fmt.Printf("%#v\n\n", c.Employee)

	// fmt.Printf("%v %v\n", c.Name, c.Empid)
	// fmt.Printf("%v %v\n", c.Emp.Name, c.Emp.Empid)
	// fmt.Printf("%v %v\n", c.Employee.Name, c.Employee.Empid)
}
