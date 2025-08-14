// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案。

package main

import "fmt"

// 暴力解法
// 两层for循环，遍历所有可能的组合
// 检查 nums[i] + nums[j] 是否等于 target，如果等于则返回 i 和 j
// 如果遍历完数组，没有找到符合条件的元素，则返回空数组
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 哈希表解法
// 遍历数组，将每个元素的值和索引存入哈希表
// 遍历数组，计算目标值与当前元素的差值，如果差值在哈希表中存在，则返回差值的索引和当前元素的索引
// 如果差值不在哈希表中，则将当前元素的值和索引存入哈希表, 数值：索引
// 如果遍历完数组，没有找到符合条件的元素，则返回空数组
// nums = [2,7,11,15] target = 18
// complement = 18 - 2 = 16
// 遍历到2时，complement = 18 - 2 = 16 不满足，Map: 2:0
// 遍历到7时，complement = 18 - 7 = 11 不满足，Map: 2:0 7:1
// 遍历到11时，complement = 18 - 11 = 7 满足，返回Map[7] = 1 和当前索引2
// 遍历到15时，complement = 18 - 15 = 3 不满足，Map: 2:0 7:1 11:2 15:3
// 遍历到18时，complement = 18 - 18 = 0 不满足，Map: 2:0 7:1 11:2 15:3 18:4
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSumHash(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := hashMap[complement]; ok {
			return []int{hashMap[complement], i}
		}
		hashMap[nums[i]] = i
		fmt.Println(hashMap)
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumHash(nums, target))
}
