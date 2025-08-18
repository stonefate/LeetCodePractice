// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。

package main

import "fmt"

// 计算公式：min(height[i], height[j]) * (j - i)
// 双指针，从两端向中间移动，每次移动较小的那个指针
// 计算当前的面积，更新最大面积
// 移动指针，移动较小的那个指针
// 返回最大面积
// height = [left=1,8,6,2,5,4,8,3,right=7]
// left = 0, right = 8
// start area = min(1, 7) * (8 - 0) = 8, maxArea = 8
// height[left] < height[right], left++
// height = [1,left=8,6,2,5,4,8,3,right=7]
// left = 1, right = 8
// area = min(8, 7) * (8 - 1) = 49, maxArea = 49
// height[left] > height[right], right--
// height = [1,left=8,6,2,5,4,8,right=3,7]
// left = 1, right = 3
// area = min(8, 3) * (7 - 1) = 18, maxArea = 18
// height[left] > height[right], right--
// ....
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
