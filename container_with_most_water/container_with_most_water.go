//Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.

package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var maxPossibleHeight, testVolume, maxVolume, leftIndex, rightIndex int
	maxVolume, leftIndex, rightIndex = 0, 0, len(height) - 1
	for leftIndex < rightIndex {
		if height[leftIndex] <= height[rightIndex] {
			maxPossibleHeight = height[leftIndex]
		} else {
			maxPossibleHeight = height[rightIndex]
		}

		testVolume = maxPossibleHeight * (rightIndex - leftIndex)
		if testVolume > maxVolume {
			maxVolume = testVolume
		}

		if height[leftIndex] <= height[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}
	return maxVolume
}

func main() {
	print(maxArea([]int{1,1}))
}
