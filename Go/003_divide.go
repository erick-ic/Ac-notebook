// Difficulty: Easy
// Status: Done
// Notes: 两数相除(藤蔓面试)
package main

import (
	"errors"
	"fmt"
)

// Divide 计算两数相除，返回结果和错误
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func main() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	result, err = Divide(5, 0)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}
}
