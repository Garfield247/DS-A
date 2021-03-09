// Package main provides ...
package main

import "fmt"

func main() {
	var nums = []int{6, 4, 7, 2, 9, 3, 1, 5, 8}
	fmt.Println(nums)
	sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func sort(array []int, leftPoint int, rightBound int) {

	if leftPoint == rightBound {
		fmt.Println("只剩一个元素了")
		return
	}
	//确定中点
	mid := leftPoint + (rightBound-leftPoint)/2
	// 左边排序
	sort(array, leftPoint, mid)
	//右边排序
	sort(array, mid+1, rightBound)

	merge(array, leftPoint, mid+1, rightBound)

}

func merge(array []int, leftPoint int, rightPoint int, rightBound int) {
	var mid int = rightPoint - 1
	temp := make([]int, rightBound-leftPoint+1)
	var tempPoint int = 0

	var i int = leftPoint
	var j int = rightPoint

	for i <= mid && j <= rightBound {
		fmt.Printf("1|i:%d|j:%d|tempPoint:%d|temp:%v|array:%v\n", i, j, tempPoint, temp, array)

		if array[i] <= array[j] {
			fmt.Printf("array[%d]:%d\n", i, array[i])
			temp[tempPoint] = array[i]
			fmt.Println("temp:", temp)
			i++
			tempPoint++
		} else {
			temp[tempPoint] = array[j]
			fmt.Printf("array[%d]:%d\n", j, array[j])
			fmt.Println("temp:", temp)
			j++
			tempPoint++
		}
	}
	for i <= mid {
		temp[tempPoint] = array[i]
		tempPoint++
		i++
	}
	for j <= rightBound {
		temp[tempPoint] = array[j]
		tempPoint++
		j++

	}

	for x := 0; x < len(temp); x++ {
		array[x+leftPoint] = temp[x]
	}
	fmt.Println(array)
}
