package questions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FizzBuzz(start int, to int) string {
	var result string
	for i := start; i <= to; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result += "FizzBuzz"
		case i%3 == 0:
			result += "Fizz"
		case i%5 == 0:
			result += "Buzz"
		default:
			result += fmt.Sprintf("%d", i)
		}
	}
	return result
}

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		start    int
		to       int
		expected string
	}{
		{1, 1, "1"},
		{1, 5, "12Fizz4Buzz"},
		{1, 15, "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test %d", index), func(t *testing.T) {
			result := FizzBuzz(testCase.start, testCase.to)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
