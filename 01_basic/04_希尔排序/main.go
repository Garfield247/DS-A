// Package main provides ...
package main

import (
	"fmt"
	"time"
)

var count int = 0

func main() {
	var nums = []int{213, 42, 32, 432, 12, 543, 256, 65, 856, 74, 234, 532, 43, 92, 63, 433}
	fmt.Println(nums)
	shellSort(nums)

}
func shellSort(array []int) {

	//step := 4
	//insertSort(array, step)
	t1 := time.Now()
	for step := 4; step >= 1; step /= 2 {
		fmt.Println("step:", step)
		for i := step; i < len(array); i++ {
			for j := i; j >= step; j -= step {
				if array[j] < array[j-step] {
					swap(j, j-step, array)
					count++
				}
			}
		}
		fmt.Println(array)
	}
	t2 := time.Now()
	fmt.Println("插入次数:", count)
	fmt.Println("耗时:", t2.Sub(t1))

}

func swap(index1 int, index2 int, array []int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}
