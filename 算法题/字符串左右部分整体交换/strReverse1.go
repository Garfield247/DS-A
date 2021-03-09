// Package main provides ...
package main

import "fmt"

func main() {
	str := "abcdef"
	leftsize := 4
	str = StrReverse(str, leftsize)
	fmt.Println(str)
}

func StrReverse(str string, leftsize int) string {
	s := []rune(str)
	fmt.Println(s)
	length := len(s)

	// 左边逆序
	reverse(s, 0, leftsize-1)
	// 右边逆序
	reverse(s, leftsize, length-1)
	// 整体逆序
	reverse(s, 0, length-1)
	return string(s)
}

func reverse(str []rune, start, stop int) {

	for x, y := start, stop; x < y; x, y = x+1, y-1 {
		temp := str[x]
		str[x] = str[y]
		str[y] = temp
	}
	fmt.Println(str)
}
