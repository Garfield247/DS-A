// Package main provides ...
package main

import (
	"fmt"
	"time"
)

// 选择排序

func main() {
	Array1 := [9]int{6, 4, 7, 2, 9, 3, 1, 5, 8}
	fmt.Println("Array1:", Array1)
	t1 := time.Now()
	for i := 0; i < len(Array1)-1; i++ {
		minPos := i
		// 从当前位置开始找出最小值的索引
		for j := i + 1; j < len(Array1); j++ {
			if Array1[j] < Array1[minPos] {
				minPos = j
			}
		}
		fmt.Println("本次minPos:", minPos)
		// 将最小值放到最前
		tmp := Array1[i]
		Array1[i] = Array1[minPos]
		Array1[minPos] = tmp

		fmt.Printf("Array1:%v\n", Array1)
	}
	t2 := time.Now()

	fmt.Printf("最终结果:%v\n耗时:%v", Array1, t2.Sub(t1))
}
