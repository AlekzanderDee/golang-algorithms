//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLength := 1
	curMaxLength := 1
	var startPos int = 0
	var charsBitMap = map[byte]bool{}

	for startPos <= len(s)-2 {
		charsBitMap[s[startPos]] = true

		for i := startPos + 1; i < len(s); i++ {
			if charsBitMap[s[i]] == false {
				curMaxLength++
				charsBitMap[s[i]] = true
				if curMaxLength > maxLength {
					maxLength = curMaxLength
				}
			} else {
				curMaxLength = 1
				charsBitMap = map[byte]bool{}
				if curMaxLength > maxLength {
					maxLength = curMaxLength
				}
				break
			}
		}
		startPos++
	}
	return maxLength
}

func main() {
	type TestCase struct {
		InpString      string
		ExpectedResult int
	}

	tests := []TestCase{
		{
			InpString:      "abcabcbb",
			ExpectedResult: 3,
		},
		{
			InpString:      "bbbbb",
			ExpectedResult: 1,
		},
		{
			InpString:      "pwwkew",
			ExpectedResult: 3,
		},
		{
			InpString:      "dvdf",
			ExpectedResult: 3,
		},
		{
			InpString:      "au",
			ExpectedResult: 2,
		},
	}

	for _, test := range tests {
		res := lengthOfLongestSubstring(test.InpString)
		if res != test.ExpectedResult {
			fmt.Printf("InputStr=%s; Expected result=%v; Fact result=%v\n",
				test.InpString, test.ExpectedResult, res)
		}
	}
	fmt.Print("completed")
}
