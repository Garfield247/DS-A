package main

import (
	"fmt"
	"time"
)

func main() {
	//nums := []int{3, 9, 7, 1, 2, 5, 8, 4, 6}
	nums := []int{72, 6, 57, 88, 60, 42, 83, 73, 48, 85}
	t1 := time.Now()
	quickSort(nums, 0, len(nums)-1)
	t2 := time.Now()
	fmt.Printf("排序结果:%v,耗时啊:%v \n", nums, t2.Sub(t1))

}

func quickSort(array []int, start int, stop int) {
	if stop < start {
		return
	}
	var i int = start
	var j int = stop
	var x int = array[start]

	for i < j {
		for i < j && array[j] > x {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}

		for i < j && array[i] < x {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	if i == j {

		array[i] = x
	}
	fmt.Println(array)
	quickSort(array, start, i-1)
	quickSort(array, i+1, stop)
}
