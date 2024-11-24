package main

import (
	"fmt"
	"github.com/arthurasanaliev/mental-math-go/pkg/questions"
	"time"
)

// main is the entry point of the program
func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	practice := true

	for practice {
		fmt.Println("Choose a topic you want to practice:")
		fmt.Println("1 -- Addition")
		fmt.Println("2 -- Substraction")
		fmt.Println("3 -- Multiplication")
		fmt.Println("4 -- Division")

		var questionType int
		fmt.Scan(&questionType)

		switch questionType {
		case 1:
			questions.Addition()
		case 2:
			questions.Substraction()
		case 3:
			questions.Multiplication()
		case 4:
			questions.Division()
		default:
			practice = false
		}
	}
}
