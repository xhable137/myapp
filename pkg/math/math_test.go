package math_test

import (
	"testing"

	"github.com/xhable137/myapp/pkg/math"
)

func TestAverage(t *testing.T) {
	if _, err := math.Average(nil); err == nil {
		t.Error("Expected error for empty slice, got nil")
	}
	avg, err := math.Average([]float64{2, 4, 6})
	if err != nil || avg != 4.0 {
		t.Errorf("Average([]float64{2,4,6}) = %v, %v; want 4.0, nil", avg, err)
	}
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11}
	for _, p := range primes {
		if !math.IsPrime(p) {
			t.Errorf("IsPrime(%d) = false; want true", p)
		}
	}
	nonPrimes := []int{-1, 0, 1, 4, 6}
	for _, n := range nonPrimes {
		if math.IsPrime(n) {
			t.Errorf("IsPrime(%d) = true; want false", n)
		}
	}
}

func TestMaxIntSlice(t *testing.T) {
	if _, err := math.MaxIntSlice(nil); err == nil {
		t.Error("Expected error for empty slice, got nil")
	}
	max, err := math.MaxIntSlice([]int{1, 3, 2})
	if err != nil || max != 3 {
		t.Errorf("MaxIntSlice([]int{1,3,2}) = %d, %v; want 3, nil", max, err)
	}
}
