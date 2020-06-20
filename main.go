package main

import (
	"fmt"
	"sicp/chapter1"
)

func main() {
	fmt.Print("Hello World \n")
	fmt.Println(chapter1.Factorial(5))
	fmt.Println(chapter1.RecursiveFibonacci(50))
	fmt.Println(chapter1.CoinChange(75))
	fmt.Println(chapter1.Fibonacci3(20))
	fmt.Println(chapter1.Sine(12.15))
	fmt.Println(chapter1.Pow(2, 10))
}
