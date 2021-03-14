package main

import (
	"fmt"
	"math/rand"
	"time"
)

func basic() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("sum from 1 to 100 is:", sum)
}

func asWhile() {
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}

func asWhileTrue() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Println("Writting inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finished !")
			break
		}
		fmt.Println(num)
	}
}

func withContinue() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}

func main() {
	basic()
	asWhile()
	asWhileTrue()
	withContinue()
}
