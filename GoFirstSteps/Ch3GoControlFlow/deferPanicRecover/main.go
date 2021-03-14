package main

import (
	"fmt"
	"io"
	"os"
)

func deferExample() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred: ", i)
		fmt.Println("regular: ", i)
	}
}

func deferClose() {
	f, err := os.Create("notes.txt")
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.WriteString(f, "learning go")
	if err != nil {
		return
	}

	f.Sync()
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicing")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("defer in g(): ", i)
	fmt.Println("println in g(): ", i)
	g(i + 1)
}

func main() {
	deferExample()
	fmt.Println()
	deferClose()
	fmt.Println()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in main: ", r)
		}
	}()
	g(0)
	fmt.Println("this will not print")
}
