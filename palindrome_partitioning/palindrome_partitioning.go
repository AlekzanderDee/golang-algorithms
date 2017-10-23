//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return all possible palindrome partitioning of s.
//
//For example, given s = "aab",
//Return
//
//[
//["aa","b"],
//["a","a","b"]
//]


package main

import "fmt"

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

func partition(s string) [][]string {
	res := [][]string{}
	if isPalindrome(s) {
		res = append(res, []string{s})
	}
	strLen := len(s)
	for pos := 1; pos < strLen; pos++ {
		leftStr := s[:pos]
		if isPalindrome(leftStr) {
			rightStr := s[pos:]
			for _, palindrome := range partition(rightStr) {
				s := []string{leftStr, }
				s = append(s, palindrome...)
				res = append(res, s)
			}
		}
	}
	return res
}

func main(){
	fmt.Printf("%v\n", partition("aabbcc"))
	fmt.Printf("%v\n", partition("bb"))
	fmt.Printf("%v\n", partition("bbbb"))
	fmt.Printf("%v\n", partition("abb"))
}
