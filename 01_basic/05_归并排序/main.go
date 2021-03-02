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

	for leftPoint < mid && rightPoint < rightBound {
		fmt.Printf("1|leftPoint:%d|rightPoint:%d|tempPoint:%d|temp:%v|array:%v\n", leftPoint, rightPoint, tempPoint, temp, array)

		if array[leftPoint] <= array[rightPoint] {
			fmt.Printf("array[%d]:%d\n", leftPoint, array[leftPoint])
			temp[tempPoint] = array[leftPoint]
			fmt.Println("temp:", temp)
			leftPoint++
			tempPoint++
		} else {
			temp[tempPoint] = array[rightPoint]
			fmt.Printf("array[%d]:%d\n", rightPoint, array[rightPoint])
			fmt.Println("temp:", temp)
			rightPoint++
			tempPoint++
		}
	}
	for leftPoint <= mid {
		temp[tempPoint] = array[leftPoint]
		tempPoint++
		leftPoint++
	}
	for rightPoint < rightBound {
		temp[tempPoint] = array[rightPoint]
		tempPoint++
		rightPoint++

	}

	for i := leftPoint; i <= rightBound; i++ {
		array[i] = temp[i-leftPoint]
	}
	fmt.Println(array)
}
