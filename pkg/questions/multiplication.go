package questions

import "fmt"

// Multiplication implements multiplication practice
func Multiplication() {
	err := base("*", 60)
	if err != nil {
		fmt.Println(err)
	}
}
