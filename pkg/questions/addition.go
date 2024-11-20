package questions

import "fmt"

// Addition implements addition practice
func Addition() {
	err := base("+", 30)
	if err != nil {
		fmt.Println(err)
	}
}
