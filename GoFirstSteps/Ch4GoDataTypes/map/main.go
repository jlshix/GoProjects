package main

import "fmt"

func mInit() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)
}

func mAppend() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)
}

func mAccess() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println("Christy's age is", studentsAge["christy"])
}

func mAccessExist() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
}

func mDelete() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	delete(studentsAge, "john")
	delete(studentsAge, "christy")
	fmt.Println(studentsAge)
}

func mForLoop() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
	for _, age := range studentsAge {
		fmt.Printf("Age is %d\n", age)
	}
	for name := range studentsAge {
		fmt.Printf("Name is %s\n", name)
	}
}

func main() {
	mInit()
	mAppend()
	mAccess()
	mAccessExist()
	mDelete()
	mForLoop()
}
