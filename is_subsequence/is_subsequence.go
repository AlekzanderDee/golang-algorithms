/*
Difficulty:Medium

Given a string s and a string t, check if s is subsequence of t.

You may assume that there is only lower case English letters in both s and t.
t is potentially a very long (length ~= 500,000) string, and s is a short string (<=100).

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
s = "abc", t = "ahbgdc"

Return true.

Example 2:
s = "axc", t = "ahbgdc"

Return false.

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?
*/


package main

import (
	"reflect"
	"fmt"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
		panic("Assertion error\n")
	}
}


func isSubsequence(s string, t string) bool {
	posS := 0
	posT := 0
	lenT := len(t)
	lenS := len(s)

	for posS < lenS && posT < lenT {
		if t[posT] == s[posS] {
			posS++
		}
		posT++
	}
	return posS == lenS
}


func main(){
	type TestCase struct {
		S string
		T string
		ExpectedResult bool
	}

	tests := []TestCase {
		{
			S: "abc",
			T: "ahbgdc",
			ExpectedResult: true,
		},
		{
			S: "axc",
			T: "ahbgdc",
			ExpectedResult: false,
		},
	}

	for _, test := range tests {
		res := isSubsequence(test.S, test.T)
		assertEq(test.ExpectedResult, res)
	}

	fmt.Print("Completed")

}

