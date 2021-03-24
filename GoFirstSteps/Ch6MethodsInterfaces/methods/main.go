package main

import (
	"fmt"
	"geometry"
	"strings"
)

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

// ---

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

// ---

type coloredTriangle struct {
	triangle
	color string
}

func (t coloredTriangle) perimeter() int {
	return t.triangle.perimeter() * 2
}

func main() {

	t := triangle{3}
	fmt.Println(t.perimeter())
	t.doubleSize()
	fmt.Println(t.perimeter())

	// ---

	s := upperstring("Learing Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())

	// ---

	ct := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("ct Size:", ct.size)
	fmt.Println("ct Perimeter", ct.perimeter())

	// ---

	gt := geometry.Triangle{}
	gt.SetSize(3)
	fmt.Println("gt Perimeter", gt.Perimeter())

	// gt.size undefined (cannot refer to unexported field or method size)
	// fmt.Println("size: ", gt.size)
}
