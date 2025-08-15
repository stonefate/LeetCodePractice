// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

package main

import "fmt"

// 遍历数组，对于每个数字，我们试图以它为起点，不断向上查找（x+1, x+2, ...）是否存在于数组中，从而确定序列的长度
// nums = [100,4,200,1,3,2]
// 100, 100-1 不存在，是起点，100+1是否存在，不存在，100的序列长度是1
// 4, 4-1 存在， 不是起点，跳过
// 200，200-1 不存在，是起点，200+1是否存在，不存在，200的序列长度是1
// 1, 1-1 不存在，是起点，1+1是否存在，存在，1+2是否存在，存在，1+3是否存在，存在，1+4不存在，1的序列长度是4
// 3, 3-1 存在，不是起点，跳过
// 2, 2-1 存在，不是起点，跳过
func longestConsecutive(nums []int) int {
	// 将数组转换为map，key为数字，value为true
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	fmt.Println(hashMap)

	// 遍历数组，对于每个数字，我们试图以它为起点，不断向上查找（x+1, x+2, ...）是否存在于数组中，从而确定序列的长度
	longestStreak := 0
	for _, num := range nums {
		fmt.Println(num, "num")
		if hashMap[num-1] {
			fmt.Println(num, "不是起点")
			continue
		}
		// 如果当前数字是起点，则开始计算序列长度
		currentNum := num
		currentStreak := 1
		fmt.Println(currentNum, "是起点")
		for hashMap[currentNum+1] {
			fmt.Println(currentNum, "继续查找下一个")
			currentNum++
			currentStreak++
		}
		// 更新最长序列长度
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
		fmt.Println(longestStreak, "longestStreak")
	}
	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("max length", longestConsecutive(nums))
}
