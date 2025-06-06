// Difficulty: Easy
// Status: Done
// Notes: 有序数组的平方(977)
package main

import "fmt"

/*
给你一个按非递减顺序排序的整数数组nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序?

非递减顺序:每一个元素都小于或等于它后面的元素。
时间复杂度：O(n)，遍历了一次数组。
空间复杂度：O(n)，用了一个额外结果数组。
*/

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]
		if leftSq > rightSq {
			res[pos] = leftSq
			left++
		} else {
			res[pos] = rightSq
			right--
		}
		pos--
	}
	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	fmt.Println("输入数组：", nums)
	fmt.Println("平方排序结果：", result)
}
