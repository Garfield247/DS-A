// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	var nums = []int{6, 4, 7, 2, 9, 3, 1, 5, 8}
	fmt.Println(nums)
	t1 := time.Now()
	length := len(nums)

	for i := 1; i < length; i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				left(nums, j)
				fmt.Println(nums)
			}
		}
	}
	t2 := time.Now()
	fmt.Printf("排序结果:%v,耗时:%v\n", nums, t2.Sub(t1))

}

func left(nums []int, index int) {
	tmp := nums[index]
	nums[index] = nums[index-1]
	nums[index-1] = tmp

}
