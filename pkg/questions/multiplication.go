package questions

import "fmt"

// Multiplication implements multiplication practice
func Multiplication() {
	err := base("*")
	if err != nil {
		fmt.Println(err)
	}
}
