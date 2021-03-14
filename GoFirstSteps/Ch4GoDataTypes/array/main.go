package main

import "fmt"

func declare() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}

func aInit() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}

func ellipses1() {
	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}

func ellipses2() {
	numbers := [...]int{99: -1, 100: -2}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[100])
	fmt.Println("Length:", len(numbers))
}

func twoDimension() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}

func main() {
	declare()
	aInit()
	ellipses1()
	ellipses2()
	twoDimension()
}
