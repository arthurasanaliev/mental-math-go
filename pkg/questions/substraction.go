package questions

import "fmt"

// Substraction implements substraction practice
func Substraction() {
	err := base("-")
	if err != nil {
		fmt.Println(err)
	}
}
