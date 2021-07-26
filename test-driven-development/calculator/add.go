package calculator

import "fmt"

func Add(a, b int) int {
	if a == 6 {
		fmt.Println("tidak tercover")
	}
	return a + b
}
