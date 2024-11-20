package questions

import "fmt"

// TODO: add numbers with more digits

// Division implements division practice
func Division() {
	err := base("/", 30)
	if err != nil {
		fmt.Println(err)
	}
}
