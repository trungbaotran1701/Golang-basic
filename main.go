package main

import (
	"fmt"
	"math"
)

// 1. Create a function to reverse a string.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 2. Create a function to check if a string is a palindrome.
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 3. Create a function to calculate the factorial of a given number.
func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

// 4. Create a function to find the greatest common divisor of two numbers. => can find least common multiple of two numbers.
func gcd(a int, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

// 6. Create a function to check if a given number is a prime number.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 7. Create a function to create a Fibonacci series up to a given number.
func createFibonacciSeries(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// 8. Create a function to bubble sort an array of numbers.
func bubbleSort(arr []int) {
	lengthOfArr := len(arr)
	for i := 0; i < lengthOfArr-1; i++ {
		for j := 0; j < lengthOfArr-1-j; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 9. Create a function to calculate the sum of an array of numbers.
func calculateSum(arr []int) int {
	sum := 0
	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]
	}
	return sum
}

// 10. Create a function to find the maximum element in an array of numbers
func findMaxNumber(arr []int) int {
	maxNumber := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxNumber < arr[i] {
			maxNumber = arr[i]
		}
	}
	return maxNumber
}

func main() {

	//Exercise 1 and 2
	str := "Hello Hectagon Finance"
	fmt.Println()
	fmt.Println("Original string:", str)
	fmt.Println()
	fmt.Println("Ex1: Reversed string:", reverse(str))
	fmt.Println()
	fmt.Println("Ex2: Is Palindrome:", isPalindrome(str))
	fmt.Println()

	//Exercise 3
	x := 6
	fmt.Println(fmt.Sprintf("Ex3: Factorial of %d is: %d", x, calculateFactorial(x)))
	fmt.Println()

	//Exercise 4 and 5
	x1 := 4
	x2 := 6
	result := gcd(x1, x2)
	if result == 0 {
		fmt.Println("Ex4-Ex5: The GCD and LCM of 0 and 0 is undefined.")
	} else {
		fmt.Println(fmt.Sprintf("Ex4: The greatest common divisor of %d and %d is: %d", x1, x2, result))
		fmt.Println()
		if result == x1 || result == x2 {
			fmt.Println(fmt.Sprintf("Ex5: The least common multiple of %d and %d is undefined", x1, x2))
		} else {
			fmt.Println(fmt.Sprintf("Ex5: The least common multiple of %d and %d is: %d", x1, x2, x1*x2/result))
		}
	}
	fmt.Println()

	//Exercise 6
	x6 := 6
	result6 := isPrime(x6)
	if result6 {
		fmt.Println(fmt.Sprintf("Ex6: %d is a prime number", x6))
	} else {
		fmt.Println(fmt.Sprintf("Ex6: %d is not a prime number", x6))
	}
	fmt.Println()

	//Exercise 7
	x7 := 5
	if x7 < 3 {
		fmt.Println("You can not create Fibonacci series with this number")
	} else {
		result7 := createFibonacciSeries(x7)
		fmt.Println(fmt.Sprintf("Ex7: The Fibonacci series was created with %d numbers is: %d", x7, result7))
	}
	fmt.Println()

	//Exercise 8 and 9 and 10
	arr := []int{1, 4, 2, 5, 8}
	bubbleSort(arr)
	fmt.Println("Ex8: The sequence of numbers arranged in ascending order is:", arr)
	fmt.Println()
	fmt.Println("Ex9: The sum of array is:", calculateSum(arr))
	fmt.Println()
	fmt.Println("Ex10: Max number is:", findMaxNumber(arr))
}
