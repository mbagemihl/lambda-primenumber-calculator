package calc

import "math"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func GetPrimeNumbersBelowMax(max int) []int{
	var result []int
	for i :=1; i < max; i++ {
		if prime := IsPrime(i); prime == true {
			result = append(result, i)
		}
	}
	return result
}