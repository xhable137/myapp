package math

import "errors"

// ErrEmptySlice указывает, что срез пуст.
var ErrEmptySlice = errors.New("slice is empty")

// Average возвращает среднее арифметическое среза nums или ошибку, если срез пуст.
func Average(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums)), nil
}

// IsPrime возвращает true, если n — простое число.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// MaxIntSlice возвращает максимальное значение в срезе nums или ошибку, если срез пуст.
func MaxIntSlice(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}
