package main

import (
	"encoding/json"
	"fmt"
)

func declareAndInit() {
	type Employee struct {
		ID        int
		FirstName string
		LastName  string
		Address   string
	}
	var john Employee
	employee := Employee{1001, "John", "Snow", "Snow Street"}
	employee1 := Employee{LastName: "Snow", FirstName: "John"}
	employee.ID = 1002
	employeeCopy := &employee
	employeeCopy.FirstName = "David"

	fmt.Println("john =", john)
	fmt.Println("employee =", employee)
	fmt.Println("employee1 =", employee1)
	fmt.Println("employeeCopy =", employeeCopy)
}

func embed1() {
	type Person struct {
		ID        int
		FirstName string
		LastName  string
		Address   string
	}

	type Employee struct {
		Information Person
		ManagerID   int
	}

	var employee Employee
	employee.Information.FirstName = "John"
	fmt.Println(employee)
}

func embed2() {
	type Person struct {
		ID        int
		FirstName string
		LastName  string
		Address   string
	}

	type Employee struct {
		Person
		ManagerID int
	}

	type Contractor struct {
		Person
		CompanyID int
	}

	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Snow"
	fmt.Println(employee.LastName)
}

func marshalUnmarshal() {
	type Person struct {
		ID        int
		FirstName string `json:"name"`
		LastName  string
		Address   string `json:"address,omitempty"`
	}

	type Employee struct {
		Person
		ManagerID int
	}

	type Contractor struct {
		Person
		CompanyID int
	}

	employees := []Employee{
		{
			Person: Person{
				LastName: "Snow", FirstName: "John",
			},
		},
		{
			Person: Person{
				LastName: "North", FirstName: "King",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Println(decoded)
}

func main() {
	declareAndInit()
	embed1()
	embed2()
	marshalUnmarshal()
}
