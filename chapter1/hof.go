package chapter1

// Square returns the square of number
func Square(x int) int {
	return x * x
}

// SquareRoot returns square root to a particular confidence.
func SquareRoot(x float64) float64 {
	return squareRootItr(1, x)
}

// Factorial returns factorial of the number
func Factorial(number int) int {
	return linearFactorial(1, number)
}

// RecursiveFibonacci get fibonacci number in inefficient way
func RecursiveFibonacci(number int) int {
	return fibn(number)
}

// CoinChange returns change of coins
func CoinChange(cash int) int {
	return counter(5, 0, cash)
}

// Sine of value x Exercise 1.14
func Sine(radian float64) float64 {
	if radian < 0.001 && radian > -0.001 {
		return radian
	}
	return 3*Sine(radian/3) - 4*cubeFloat(Sine(radian/3))
}

// Pow calculates the power to number in exponential order Section 1.4
func Pow(base int, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 {
		return Pow(base*base, exp/2)
	}
	return base * Pow(base, exp-1)
}
