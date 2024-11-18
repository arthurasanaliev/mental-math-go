package questions

import "fmt"

// Addition implements addition practice
func Addition() {
	err := base("+")
	if err != nil {
		fmt.Println(err)
	}
}
