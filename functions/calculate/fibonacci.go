package calculate

import (
	"math/big"
)

// CalculateFibonacci calculates the Fibonacci number for the given input.
// @param n query int true "Number of Fibonacci sequence (must be a positive integer)" format(bigint) minimum(1)
// @return result string "Result of Fibonacci calculation (string representation of the Fibonacci number)" format(int64)
func CalculateFibonacci(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	prev1 := big.NewInt(0)
	prev2 := big.NewInt(1)
	fib := big.NewInt(0)

	for i := uint(2); i <= n; i++ {
		fib.Add(prev1, prev2)
		prev1.Set(prev2)
		prev2.Set(fib)
	}

	return fib
}
