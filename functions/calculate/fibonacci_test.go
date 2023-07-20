package calculate_test

import (
	"testing"

	"github.com/daan0220/fib-api/functions/calculate"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFibonacci(t *testing.T) {
	tests := []struct {
		n        uint
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{2, "1"},
		{3, "2"},
		{4, "3"},
		{5, "5"},
		{10, "55"},
		{20, "6765"},
		{30, "832040"},
	}

	for _, test := range tests {
		result := calculate.CalculateFibonacci(test.n)
		assert.Equal(t, test.expected, result.String())
	}
}
