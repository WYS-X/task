package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//只出现一次的数字
	nums := []int{3, 4, 5, 6, 4, 5, 6}
	n := singleNumber(nums)
	fmt.Printf("%d 只出现了一次\n", n)

	//判断回文数
	palindrome := 234321
	if isPalindrome(palindrome) {
		fmt.Println("字符串方式：", palindrome, "是回文数")
	} else {
		fmt.Println("字符串方式：", palindrome, "不是回文数")
	}
	if isPalindrome1(palindrome) {
		fmt.Println("非字符串方式：", palindrome, "是回文数")
	} else {
		fmt.Println("非字符串方式：", palindrome, "不是回文数")
	}

	//20.有效的括号
	str1, str2, str3 := "([)]", "()", "()[]{}"
	if isValid(str1) {
		fmt.Println(str1, ": true")
	} else {
		fmt.Println(str1, ": false")
	}
	if isValid(str2) {
		fmt.Println(str2, ": true")
	} else {
		fmt.Println(str2, ": false")
	}
	if isValid(str3) {
		fmt.Println(str3, ": true")
	} else {
		fmt.Println(str3, ": false")
	}

	//14. 最长公共前缀
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println("最长公共前缀是:", prefix)

	//66. 加一
	fmt.Println("[9,9,9,9], 加1：", plusOne([]int{9, 9, 9, 9}))
	fmt.Println("[1,2,3], 加1：", plusOne([]int{1, 2, 3}))
	fmt.Println("[8,9,9,9], 加1：", plusOne([]int{8, 9, 9, 9}))

	//26. 删除有序数组中的重复项
	dupNums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	noDupNum := removeDuplicates(dupNums)
	fmt.Println("不重复的数字数量：", noDupNum)
	fmt.Println("去重后数组", dupNums)

	//56.合并区间
	mergeNums := [][]int{{4, 7}, {1, 4}}
	mergeRes := merge(mergeNums)
	fmt.Println(mergeNums, "合并后：", mergeRes)

	//1.两数之和
	ts := twoSum([]int{3, 3}, 6)
	fmt.Println("[3,3]之和：", ts)
}

// 获取只出现一次的数字
func singleNumber(nums []int) int {
	counter := make(map[int]int, len(nums))
	for _, n := range nums {
		_, isOk := counter[n]
		if isOk {
			counter[n] += 1
		} else {
			counter[n] = 1
		}
	}
	for k, v := range counter {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 判断是否回文数
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
func isPalindrome1(x int) bool {
	n := 0
	for x > 0 {
		m := x % 10
		x = x / 10
		n = n*10 + m
		if x == n || x == n/10 {
			return true
		}

	}
	return false
}

// 20.有效的括号
func isValid(s string) bool {
	for {
		t := s

		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
		if t == s && s != "" {
			return false
		} else if s == "" {
			return true
		}

	}
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	l := 0
	for _, a := range strs {
		if a == "" {
			return ""
		}
		if l == 0 || l > len(a) {
			l = len(a)
		}
	}
	prefix := make([]rune, 0, l)
	for i := 0; i < l; i++ {
		var p rune
		for _, a := range strs {
			r := []rune(a)
			if p == 0 {
				p = r[i]
			} else if p != r[i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, p)
	}
	return string(prefix)
}

// 66. 加一
func plusOne(digits []int) []int {
	l := len(digits)
	n := 0
	add := 1
	for i := l - 1; i >= 0; i-- {
		if add == 1 && i < l-1 {
			add = 0
		}
		if digits[i]+add+n < 10 {
			digits[i] += add + n
			n = 0
			break
		} else {
			digits[i] = 0
			n = 1
		}
	}
	if n == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits[:])
		digits[0] = 1
	}
	return digits
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	count := 1
	max := nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num > max {
			nums[count] = num
			max = num
			count++
		}
	}
	return count
}

// 56.合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i, sub := range intervals {
		if i == 0 {
			continue
		}
		last := res[len(res)-1]
		if sub[0] <= last[1] {
			if sub[1] > last[1] {
				last[1] = sub[1]
			}
		} else {
			res = append(res, sub)
		}
	}
	return res
}

// 1.两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for index, num := range nums {
		v, has := m[target-num]
		if has {
			return []int{index, v}
		}
		m[num] = index
	}
	return nil
}
