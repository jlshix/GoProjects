package main

import (
	"os"
	"strconv"
)

func sum(n1 string, n2 string) int {
	i1, _ := strconv.Atoi(n1)
	i2, _ := strconv.Atoi(n2)
	return i1 + i2
}

func repeat(name string) string {
	return name + name
}

func double(name *string) {
	*name = *name + *name
}

func main() {
	rv := sum(os.Args[1], os.Args[2])
	println("Sum: ", rv)
	name := os.Args[3]
	repeated := repeat(name)
	println("Repeated: ", repeated)
	double(&name)
	println("Doubled: ", name)
}
