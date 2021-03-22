package main

import (
	"fmt"
	"log"
	"os"
)

func log1() {
	log.Print("log1: Hey, I'm a log!")
}

func fatalLog() {
	log.Fatal("Hey, I'm an error log !")
	fmt.Print("Can you see me ?")
}

func panicLog() {
	log.Panic("Hey, I'm an error log !")
	fmt.Print("Can you see me ?")
}

func logToFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}

func main() {
	// log1()
	// fatalLog()
	// panicLog()
	logToFile()
}
