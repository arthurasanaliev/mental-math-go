package questions

import "fmt"

// TODO: add numbers with more digits

// Division implements division practice
func Division() {
	err := base("/")
	if err != nil {
		fmt.Println(err)
	}
}
