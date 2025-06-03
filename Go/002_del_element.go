// Difficulty: Easy
// Status: Done
// Notes: 双指针(27)
package main

import (
	"fmt"
)

/*
给你一个数组nums和一个值val，你需要原地移除所有数值等于val的元素，元素的顺序可能发生改变。
然后返回nums中与val不同的元素的数量。
*/

func removeElement(nums []int, val int) []int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	// return left
	return nums[:left]
}

// 用 left 指向下一次可以放置非 val 值的位置。
// 用遍历方式读取原数组中的所有元素 v。
// 只要不是 val，就将 v 填入 nums[left]，并移动指针。
// 这种方式：
// 时间复杂度：O(n)
// 空间复杂度：O(1)（原地操作）

func main() {
	// fmt.Println(removeElement([]int{3, 2, 4, 5}, 3)) // 2
	fmt.Println(removeElement([]int{3, 2, 4, 5}, 3)) // [2 4 5]
}
