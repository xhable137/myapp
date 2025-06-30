/*
File: cmd/myapp/main.go
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xhable137/myapp/pkg/math"
	"github.com/xhable137/myapp/pkg/models"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

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

	// Демонстрация swap
	x, y := 10, 20
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)

	// Демонстрация Person
	p := models.Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p.Greet())
	fmt.Println(p.HaveBirthday())

	// Демонстрация Person с методами
	p.Greet()
	p.HaveBirthday()
	fmt.Println(p.Age)
}
