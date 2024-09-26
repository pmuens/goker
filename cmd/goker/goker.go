package main

import (
	"fmt"

	"github.com/pmuens/goker/goker"
)

func main() {
	result := goker.Mul(42, 2)
	fmt.Printf("42 * 2 = %v\n", result)
}
