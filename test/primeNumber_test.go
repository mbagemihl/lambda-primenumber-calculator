package test

import (
	"awesomeFunction/calc"
	"testing"
)

var params = [10]bool{false, false, true, true, false, true, false, true, false, false}

func TestIsPrime(t *testing.T) {
	for i := 0; i < len(params); i++ {
		if calc.IsPrime(i) != params[i] {
			t.Fail()
		}
	}
}

func TestGetPrimeNumbersBelow9(t *testing.T) {
	var expected = []int{2,3,5,7}
	actual := calc.GetPrimeNumbersBelowMax(9)

	if len(actual)!=len(expected) {t.Fail()}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {t.Fail()}
	}
}

func BenchmarkPrimeNumbersBelow9000(b *testing.B) {
	calc.GetPrimeNumbersBelowMax(90000)
}
