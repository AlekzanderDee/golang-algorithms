/*
Difficulty:Medium

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:

Input: "babad"
Output: "bab"

Note: "aba" is also a valid answer.

Example:

Input: "cbbd"
Output: "bb"
*/


package main

import (
	"fmt"
	"reflect"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
		panic("Assertion error\n")
	}
}

func isPalindrome(s string) bool {
	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return ""
	}
	// if the whole string is a palindrome - return it
	if isPalindrome(s) {
		return s
	}

    // dynamic programming matrix
    // dp[i][j]  =
    //    true, if the substring Si…Sj is a palindrome
    //    false, otherwise

    // dp(i,j)=(dp(i+1,j−1) and S​i​​==S​j​​)
    // dp(i,i)=true - one symbol length palindrome
    // dp(i,i+1)=(S​i​​==S​i+1​​) - two same neighbors - is 2 characters long palindrome

	dp := make([][]bool, lenS)
    for row := range dp {
    	dp[row] = make([]bool, lenS)
    }
	maxLen := 0
	maxPalindromeString := ""

	for j := 0; j < lenS; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				// one character long palindrome case
				dp[i][j] = true
			} else if j == i + 1 {
				// check if neighbors are the same characters (2 characters palindrome)
				eq := s[i] == s[j]
				dp[i][j] = eq
				if eq {
					if maxLen < 2 {
						maxLen = 2
						maxPalindromeString = s[i: j + 1]
					}
				}
			} else {
				// if we have the same characters at the beginning and at the end of the [i...j] substring
				// and if [i+1..j-1] is also a palindrome then it obvious that [i...j] is palindrome as well
				//
				// Example: "wefhfe"
				// if we know that "fhf" is a palindrome then "efhfe" will be a palindrome as well, because we've
				// added the same characters to the beginning and to the end
				if i <= j && s[i] == s[j] && dp[i+1][j-1] == true {
					dp[i][j] = true
					if maxLen < j - i + 1 {
						maxLen = j - i + 1
						maxPalindromeString = s[i: j + 1]
					}
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	if maxLen == 0 {
		maxPalindromeString = s[lenS - 1:]
	}
	return maxPalindromeString
}

func main() {
	type TestCase struct {
		S string
		MaxPalindromicSubstring string
	}
	tests := []TestCase{
		{
			S: "babad",
			MaxPalindromicSubstring: "bab",
		},
		{
			S: "cbbd",
			MaxPalindromicSubstring: "bb",
		},
		{
			S: "babadada",
			MaxPalindromicSubstring: "adada",
		},
		{
			S: "b",
			MaxPalindromicSubstring: "b",
		},
		{
			S: "",
			MaxPalindromicSubstring: "",
		},
		{
			S: "abcdefedcba",
			MaxPalindromicSubstring: "abcdefedcba",
		},

	}
	for _, test := range tests {
		res := longestPalindrome(test.S)
		assertEq(test.MaxPalindromicSubstring, res)
	}

	fmt.Print("Completed")
}
