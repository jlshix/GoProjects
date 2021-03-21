package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found")
var ErrServerNotResponding = errors.New("server has failed to respond to get the employee information")

func apiCallEmployee(id int) (*Employee, error) {
	if id == 1000 {
		return &Employee{}, nil
	}
	return nil, ErrNotFound
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	if err != nil {
		return nil, fmt.Errorf("Got an error: %v", err)
	}
	return employee, nil
}

func getInformationWithRetry(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(id)
		if err == nil {
			return employee, nil
		}
		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}
	return nil, ErrServerNotResponding
}

func main() {
	employee, err := getInformationWithRetry(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("NOT FOUND ", err)
	} else if errors.Is(err, ErrServerNotResponding) {
		fmt.Println("SERVER NOT RESPONDING ", err)
	} else {
		fmt.Println(employee)
	}
}
