package questions

import (
	"fmt"
	"github.com/arthurasanaliev/mental-math-go/pkg/utils"
	"time"
)

// Addition implements addition practice
func Addition() {
	fmt.Println("Addition practice starts in ...")
	utils.CountDown()

	countCorrect := 0
	wrongs := [][]int{}

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		a, b := utils.GetNumbers(10, 99)
		fmt.Printf("%d + %d = ", a, b)
		var res int
		fmt.Scan(&res)
		if res == a+b {
			countCorrect++
		} else {
			wrongs = append(wrongs, []int{a, b, res})
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
		fmt.Printf("%d + %d = %d\n", pair[0], pair[1], pair[2])
	}
	fmt.Println()
}
