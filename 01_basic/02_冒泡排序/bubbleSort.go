// 冒泡排序
// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	var nums = [9]int{6, 4, 7, 2, 9, 3, 1, 5, 8}
	t1 := time.Now()
	// 控制轮数
	// 每轮把最大值移动到最后端
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			// 比较相邻的两个值的大小
			if nums[j] > nums[j+1] {
				//若前者大于后者则置换位置
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
		fmt.Println(nums)
	}
	t2 := time.Now()
	fmt.Printf("排序结果:%v,耗时:%v\n", nums, t2.Sub(t1))
}
