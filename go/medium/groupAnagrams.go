// 给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。
// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:
// 输入: strs = [""]
// 输出: [[""]]
// 示例 3:
// 输入: strs = ["a"]
// 输出: [["a"]]

// 把字母数量相同的字符串分到一组

package main

import (
	"fmt"
	"sort"
	"strings"
)

// 26个字母分别对应1到26，""对应0，计算每个str的字符对应的数值之和，如果数值相同，则将str分到一组
// ["eat", "tea", "tan", "ate", "nat", "bat"] [23, 23, 32, 23, 32,]
func getStrSum(str string) int {
	sum := 0
	for _, char := range str {
		sum += int(char - 'a')
	}
	fmt.Println(str, sum)
	return sum
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	hashMap := make(map[int][]string)
	for _, str := range strs {
		sum := getStrSum(str)
		hashMap[sum] = append(hashMap[sum], str) // 将sum作为key，str作为value
	}
	for _, value := range hashMap {
		result = append(result, value)
	}
	return result
}

// 排序后，如果两个字符串相同，则它们是字母异位词
// ["eat", "tea", "tan", "ate", "nat", "bat"] ["aet", "aet", "ant", "aet", "ant", "abt"]
func groupAnagramsSort(strs []string) [][]string {
	result := [][]string{}
	hashMap := make(map[string][]string)
	for _, str := range strs {
		sortStr := strings.Split(str, "")
		sort.Strings(sortStr)
		hashMap[strings.Join(sortStr, "")] = append(hashMap[strings.Join(sortStr, "")], str) // 将排序后的字符串作为key，原字符串作为value
		fmt.Println(hashMap)
	}
	for _, value := range hashMap {
		result = append(result, value)
	}
	return result
}

func countStr(str string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range str {
		countMap[char]++ // 统计每个字符出现的次数 , 'a' -> 97
	}
	fmt.Println(countMap)
	return countMap // 返回一个map，key是字符，value是字符出现的次数
}

// 计数法，统计每个字符串中每个字符出现的次数，如果两个字符串的计数相同，则它们是字母异位词
// ["eat", "tea", "tan", "ate", "nat", "bat"] [[97:1 101:1 116:1], 97:1 101:1 116:1], [97:1 110:1 116:1], [97:1 101:1 116:1], map[97:1 110:1 116:1], map[97:1 98:1 116:1]]
func groupAnagramsCount(strs []string) [][]string {
	result := [][]string{}
	hashMap := make(map[string][]string)
	for _, str := range strs {
		countMap := countStr(str)
		countMapStr := ""
		for char, count := range countMap {
			countMapStr += fmt.Sprintf("%c%d", char, count)
		}
		hashMap[countMapStr] = append(hashMap[countMapStr], str) // 将countMapStr作为key，原字符串作为value
	}
	for _, value := range hashMap {
		result = append(result, value)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("method1", groupAnagrams(strs))
	fmt.Println("method2", groupAnagramsSort(strs))
	fmt.Println("method3", groupAnagramsCount(strs))
}
