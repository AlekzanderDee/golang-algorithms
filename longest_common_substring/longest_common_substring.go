package main

import "fmt"

func longestCommonSubstring(sX, sY string) string {
	var maxLength uint8 = 0
	var newLength uint8 = 0
	var maxX int
	// initializing empty table
	stringsTable := make([][]uint8, len(sY))
	for i := range stringsTable {
		stringsTable[i] = make([]uint8, len(sX))
	}

	for x := 0; x < len(sX); x++ {
		for y := 0; y < len(sY); y++ {
			if sX[x] == sY[y] {
				if y > 0 && x > 0 {
					newLength = stringsTable[y-1][x-1] + 1
				} else {
					newLength = 1
				}
				stringsTable[y][x] = newLength
				if newLength > maxLength {
					maxLength = newLength
					maxX = x
				}
			}
		}
	}
	if maxLength > 0 {
		subStrSlice := make([]byte, maxLength)
		for ; maxLength > 0; maxLength-- {
			subStrSlice[maxLength-1] = sX[maxX]
			maxX--
		}
		return string(subStrSlice)
	}
	return ""
}

func main() {
	type TestCase struct {
		S1             string
		S2             string
		ExpectedResult string
	}

	tests := []TestCase{
		{
			S1:             "teststring1",
			S2:             "teststring1",
			ExpectedResult: "teststring1",
		},
		{
			S1:             "teststring1",
			S2:             "anotherstringtotest",
			ExpectedResult: "string",
		},
		{
			S1:             "teststring1",
			S2:             "1",
			ExpectedResult: "1",
		},
		{
			S1:             "teststring1",
			S2:             "best",
			ExpectedResult: "est",
		},
		{
			S1:             "teststring1",
			S2:             "abc",
			ExpectedResult: "",
		},
		{
			S1:             "abc",
			S2:             "",
			ExpectedResult: "",
		},
	}

	for _, test := range tests {
		res := longestCommonSubstring(test.S1, test.S2)
		if res != test.ExpectedResult {
			fmt.Printf("S1=%s; S2=%s; Expected result=%s; Fact result=%s",
				test.S1, test.S2, test.ExpectedResult, res)
		}
	}
	fmt.Print("completed")
}
