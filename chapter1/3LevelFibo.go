package chapter1

import (
	"fmt"
	"log"
	"time"
)

// Fibonacci3 returns fibonacciOf 3 Excercise 1.11
func Fibonacci3(n int) int {

	timeStart := time.Now().Nanosecond()
	testValue1 := recursive(n)
	timeEnd := time.Now().Nanosecond()
	log.Println("Total time taken is", (timeEnd - timeStart))
	testValue2 := interative3(n)

	fmt.Println("Value are equal", testValue1, testValue2)

	return testValue2
}

func cubeInt(x int) int {
	return x * x * x
}

func cubeFloat(x float64) float64 {
	return x * x * x
}

func recursive(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return recursive(n-1) + 2*recursive(n-2) + 3*recursive(n-3)
}

func interative3(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	a := 0
	b := 1
	c := 2
	d := 2

	for i := 0; i < n-2; i++ {
		d = c + 2*b + 3*a
		a = b
		b = c
		c = d
	}
	return c
}
