package chapter1

import "math"

func squareRootItr(test float64, number float64) float64 {
	if isGoodEnough(test, number) {
		return test
	}
	newGuess := improveGuess(test, number)
	return squareRootItr(newGuess, number)

}

func average(a float64, b float64) float64 {
	return (a + b) / 2
}

func improveGuess(guess float64, number float64) float64 {
	return average(guess, float64(number)/guess)
}

func isGoodEnough(guess float64, number float64) bool {
	if math.Abs(guess*guess-number) <= 0.0001 {
		return true
	}
	return false
}

func linearFactorial(number int, counter int) int {
	if counter == 0 {
		return number
	}
	return linearFactorial(number*counter, counter-1)
}

func fibn(number int) int {
	a := 0
	b := 0
	c := 1
	for i := 0; i < number; i++ {
		a = b + c
		c = b
		b = a
	}
	return b
}

func counter(denom int, coins int, cash int) int {
	if cash == 0 {
		return 1
	} else if cash < 0 || denom == 0 {
		return 0
	} else {
		return coins + counter(denom, coins, cash-denomMatcher(denom)) + counter(denom-1, coins, cash)
	}

}

func denomMatcher(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 5
	}
	if n == 3 {
		return 10
	}
	if n == 4 {
		return 25
	}
	if n == 5 {
		return 50
	}
	return 0
}
