package lib

import (
	"fmt"
)

func init() {
	fmt.Println("[go-lib] initializing ...")
}

// Returns the sum of two numbers
func Add(a int, b int) int {
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	return a - b
}
