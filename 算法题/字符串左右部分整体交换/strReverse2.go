// Package main provides ...
package main

import "fmt"

func main() {
	str := "abcdefghljklmn"
	leftsize := 7
	str = strReverse(str, leftsize)
	fmt.Println(str)
}

func strReverse(str string, leftsize int) string {
	s := []rune(str)
	x := 0
	y := len(s) - 1
	for x < leftsize && y >= leftsize {
		fmt.Printf("x:%d | y: %d  s:%s ===> ", x, y, string(s))
		if leftsize-x > y-leftsize+1 {
			partSwap(s, x, y-leftsize+1, leftsize)
			x = x + (y - leftsize + 1)
		}
		if leftsize-x < y-leftsize+1 {
			partSwap(s, x, leftsize, y-(leftsize-x)+1)
			y = y - (leftsize - x)
		}
		if leftsize-x == y-leftsize+1 {
			partSwap(s, x, leftsize-x, x+leftsize)
			x = x + (y - leftsize + 1)
			y = y - (leftsize - x)
		}

		fmt.Println(string(s))
	}
	return string(s)
}

func partSwap(s []rune, start1 int, length int, start2 int) {
	for i, j := start1, start2; i < length; i, j = i+1, j+1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}
