/*
File: cmd/myapp/main.go
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xhable137/myapp/pkg/math"
)

func main() {
	// Демонстрация Average
	nums := []float64{1.0, 2.5, 3.5}
	avg, err := math.Average(nums)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Average of %v = %.2f\n", nums, avg)

	// Демонстрация IsPrime
	n := 17
	fmt.Printf("IsPrime(%d) = %v\n", n, math.IsPrime(n))

	// Демонстрация MaxIntSlice
	ints := []int{5, 2, 9, 3}
	max, err := math.MaxIntSlice(ints)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("MaxIntSlice(%v) = %d\n", ints, max)

	// Пример switch по времени
	hour := time.Now().Hour()
	switch {
	case hour < 6:
		fmt.Println("Good night")
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
