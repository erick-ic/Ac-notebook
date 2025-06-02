// Difficulty: Easy
// Status: Done
// Notes: 二分查找(704)
package main

import "fmt"

/*
为什么不直接用 (left + right) / 2 ？
这是初学者常见写法：
	mid := (left + right) / 2
但当 left 和 right 都很大时，left + right 可能会整型溢出，即超过 int 能表示的最大值，导致程序出错或 panic。

安全写法：避免整型溢出
	mid := left + (right - left) / 2
right - left 是两数的差，肯定不会溢出
再加上 left，整体安全。

mid := left + (right - left) / 2 仍然是中间值，和 (left + right) / 2 在数学意义上是等价的，只是写法上更安全，避免整型溢出。
*/

func search(nums []int, target int) int {
	// // 线性遍历 时间复杂度：O(n) ------最坏要看完整个数组
	// for i, num := range nums {
	// 	if num == target {
	// 		return i
	// 	}
	// }
	// return -1
	// 二分查找 时间复杂度：O(log n)，明显更优。
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	// 给定一个n个元素有序的（升序）整型数组nums和一个目标值target，写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回 -1。

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
}
