package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("Add(%d, %d) => %d\n", 1, 2, Add(1, 2))
}
