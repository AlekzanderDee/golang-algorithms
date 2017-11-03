/*
Difficulty:Easy

Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
"1",
"2",
"Fizz",
"4",
"Buzz",
"Fizz",
"7",
"8",
"Fizz",
"Buzz",
"11",
"Fizz",
"13",
"14",
"FizzBuzz"
]
*/

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i<= n; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			if i % 3 != 0 {
				res = append(res, "Buzz")
			} else if i % 5 != 0 {
				res = append(res, "Fizz")
			} else {
				res = append(res, "FizzBuzz")
			}
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

func main() {
	fmt.Printf("%#v", fizzBuzz(15))
}
