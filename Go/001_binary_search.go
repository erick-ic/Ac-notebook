// Difficulty: Easy
// Status: Done
// Notes: 二分查找(704)
package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right-left)/2 + left
		num := nums[middle]
		if num == target {
			return middle
		} else if num > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	// 给定一个n个元素有序的（升序）整型数组nums和一个目标值target，写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回 -1。

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
}
