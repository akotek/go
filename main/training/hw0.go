// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package training

import "math"

//func main() {
//	// Feel free to use the main function for testing your functions
//	fmt.Println("Hello, दुनिया!")
//}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n % 15 == 0{
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0{
		return "Buzz"
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	for i:=2; i< int(math.Sqrt(float64(n))+1); i++ {
		if n % i == 0 && n != i{
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	if len(s) <= 0 {
		return false
	}
	if len(s) == 1 { return true}

	var j = len(s)-1
	for i:= 0; i < len(s)/2; i++{
		if s[i] != s[j-i] {
			return false
		}
	}
	return true
}
