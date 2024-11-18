package questions

import (
	"errors"
	"fmt"
	"github.com/arthurasanaliev/mental-math-go/pkg/utils"
	"time"
)

// TODO: adjust scoring for different operations

// base implements base functionality for arithmetics
func base(op string) error {
	var questionType string
	switch op {
	case "+":
		questionType = "Addition"
	case "-":
		questionType = "Substraction"
	case "*":
		questionType = "Multiplication"
	case "/":
		questionType = "Division"
	default:
		return errors.New("wrong operation")
	}

	fmt.Println()
	fmt.Println(questionType, "practice starts in ...")
	utils.CountDown()

	countCorrect := 0
	wrongs := [][]int{}

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		a, b := utils.GetNumbers(10, 99)
		fmt.Printf("%d %s %d = ", a, op, b)
		var inp int
		fmt.Scan(&inp)

		var res int
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		}

		if inp == res {
			countCorrect++
		} else {
			wrongs = append(wrongs, []int{a, b, inp})
		}
	}

	elapsedTime := time.Since(startTime)
	accuracy := float32(countCorrect) / 10 * 100

	var color1, color2 string

	if elapsedTime.Seconds() >= 36.0 {
		color1 = utils.Red
	} else {
		color1 = utils.Green
	}

	if accuracy < 90.0 {
		color2 = utils.Red
	} else {
		color2 = utils.Green
	}

	fmt.Println()
	fmt.Printf("%sTime: %.2f secs%s\n", color1, elapsedTime.Seconds(), utils.Reset)
	fmt.Printf("%sAccuracy: %.2f%%%s\n", color2, float32(countCorrect)/10*100, utils.Reset)
	fmt.Println()

	fmt.Println("Answered incorrectly:")
	for _, pair := range wrongs {
		fmt.Printf("%d %s %d = %d\n", pair[0], op, pair[1], pair[2])
	}
	fmt.Println()

	return nil
}
