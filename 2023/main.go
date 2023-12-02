package main

import (
	"fmt"

	one "github.com/adorfman10/advent-of-code/1"
)

func main() {
	fmt.Printf("one.A(): %v\n", one.A())
	fmt.Printf("one.B(): %v\n", one.B())
	fmt.Printf("one.A() goroutine: %v\n", one.AGoRoutine())
	fmt.Printf("one.B() goroutine: %v\n", one.BGoRoutine())
}
