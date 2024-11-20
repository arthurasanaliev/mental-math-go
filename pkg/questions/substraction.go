package questions

import "fmt"

// Substraction implements substraction practice
func Substraction() {
	err := base("-", 35)
	if err != nil {
		fmt.Println(err)
	}
}
