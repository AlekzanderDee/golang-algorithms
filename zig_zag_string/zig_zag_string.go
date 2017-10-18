//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
//And then read line by line: "PAHNAPLSIIGYIR"
//
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string text, int nRows);
//
//convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

package main

import (
	"fmt"
)


func convertSlow(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}

	divider := numRows + (numRows - 2)
	resMatrix := make([][]byte, numRows)
	strLen := len(s)
	for rowNum := range resMatrix {
		resMatrix[rowNum] = make([]byte, 0, strLen)
	}

	for curPos := 0; curPos < strLen; curPos ++{
		adjustedPosition := curPos % divider
		if (adjustedPosition) < numRows {
			resMatrix[adjustedPosition] = append(resMatrix[adjustedPosition], s[curPos])
		} else {
			resMatrix[divider - adjustedPosition] = append(resMatrix[divider - adjustedPosition], s[curPos])
		}
	}

	res := []byte{}
	for _, row := range resMatrix {
		res = append(res, row...)
	}

	return string(res)
}

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}

	resMatrix := make([][]byte, numRows)
	index := 0
	step := 1

	for chPos := range s {
		resMatrix[index] = append(resMatrix[index], s[chPos])
		if index == 0 {
			step = 1
		} else if index == numRows - 1 {
			step = -1
		}

		index = index + step
	}

	res := []byte{}
	for _, row := range resMatrix {
		res = append(res, row...)
	}

	return string(res)
}


func main() {
	type TestCase struct {
		InpString string
		NumRows int
		ExpectedResult string
	}

	tests := []TestCase{
		{
			InpString: "PAYPALISHIRING",
			NumRows: 3,
			ExpectedResult: "PAHNAPLSIIGYIR",
		},
		{
			InpString: "",
			NumRows: 3,
			ExpectedResult: "",
		},
		{
			InpString: "PAY",
			NumRows: 3,
			ExpectedResult: "PAY",
		},
		{
			InpString: "PAYPALISHIRING",
			NumRows: 1,
			ExpectedResult: "PAYPALISHIRING",
		},
		{
			InpString: "PAYPALISHIRING",
			NumRows: 2,
			ExpectedResult: "PYAIHRNAPLSIIG",
		},
	}

	for _, test := range tests {
		convertedString := convert(test.InpString, test.NumRows)
		if convertedString != test.ExpectedResult {
			fmt.Printf("Expected result: %s, fact result: %s\n", test.ExpectedResult, convertedString)
		}
	}

	fmt.Print("Done")
}
