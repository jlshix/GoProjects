package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

// ---

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

// ---

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// ---

type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

// ---

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

// ---
func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	var s Shape = Square{3}
	printInformation(s)

	c := Circle{6}
	printInformation(c)

	// ---

	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)

	// ---
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	w := customWriter{}
	io.Copy(w, resp.Body)
}
