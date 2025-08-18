// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

package main

import (
	"fmt"
	"sort"
)

// 排序 + 双指针
// 排序后，遍历数组，对于每个元素，使用双指针找到另外两个元素，使得三个元素的和为0
// 剪枝，如果当前元素nums[i]大于0，则直接返回结果,由于数组是排好序的，后面的数都比它大，三数之和不可能等于 0, 因此，可以直接结束循环
// 去重，如果当前元素nums[i]和前一个元素nums[i-1]相同，则跳过
// 双指针查找，在 i 之后的部分，设置左指针 left = i + 1 和右指针 right = len(nums) - 1
// left < right 时，循环:
// 计算三个数的和 s = nums[i] + nums[left] + nums[right]
// 如果 s == 0，则将三个元素添加到结果中
// 如果 s > 0，则右指针左移, 说明右指针的值太大了，需要左移
// 如果 s < 0，则左指针右移, 说明左指针的值太小了，需要右移
// 移动 left 指针，跳过所有与当前 nums[left] 相同的后续元素
// 移动 right 指针，跳过所有与当前 nums[right] 相同的后续元素
// 最后，将 left 和 right 指针同时向内收缩，继续寻找下一组解
// nums = [-1,0,1,2,-1,-4]
// sort nums = [-4,-1,-1,0,1,2]
// i = 0, left = 1, value = -4, right = 5, value = 2, sum = -4 + (-1) + 2 = -3, sum < 0, 值太小，左指针右移
// i = 0, left = 2, value = -4, right = 5, value = 2, left 2 == left 1, 跳过
// i = 0, left = 3, value = 0, right = 5, value = 2, sum = -4 + 0 + 2 = -2, sum < 0, 值太小，左指针右移
// i = 0, left = 4, value = 1, right = 5, value = 2, sum = -4 + 1 + 2 = -1, sum < 0, 值太小，左指针右移
// i = 0, left = 5, value = 2, right = 5, value = 2, left == right, 结束循环
// i = 1, left = 2, value = -1, right = 5, value = 2, sum = -1 + (-1) + 2 = 0, sum = 0, 找到一组解，添加到结果中[-1,-1,2]
// 双指针内缩
// i = 1, left = 3, value = 2, right = 4, value = 1, left value > 0, 结束循环
// i = 2, left = 3, value = 0, right = 5, sum = 0 + 0 + 2 = 2, sum > 0, 值太大，右指针左移
// i = 2, left = 3, value = 0, right = 4, value = 1, left value > 0, 结束循环
// i = 3, left = 4, value = 1, right = 5, value = 2, left value > 0, 结束循环
// i = 4, left = 5, value = 2, right = 5, value = 2, left value > 0, 结束循环
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
			// 移动指针
			if sum > 0 {
				right--
			} else {
				left++
			}
			// 去重, 如果左指针和前一个元素相同，则左指针右移
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			// 去重, 如果右指针和后一个元素相同，则右指针左移
			for left < right && right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
