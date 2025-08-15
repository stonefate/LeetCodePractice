// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

package main

import "fmt"

// 反向思维, 寻找非0元素，并插入到数组中
// [0,1,0,3,12]
func moveZeroes(nums []int) {
	// 遍历数组，将非零元素插入到数组中
	nonZeroIndex := 0
	for _, num := range nums {
		fmt.Println("num", num)
		if num != 0 {
			nums[nonZeroIndex] = num
			nonZeroIndex++
			fmt.Println("nums", nums, "nonZeroIndex", nonZeroIndex)
		}
	}
	// 将非零元素后面的元素设置为0
	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
		// fmt.Println("nums", nums)
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println("nums", nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
