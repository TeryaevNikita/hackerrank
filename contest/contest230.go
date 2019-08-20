package main1

import (
	"fmt"
	"strconv"
	"strings"
)

func fibonacciWords(s1 string, s2 string, n uint64) string {

	l1, l2 := uint64(len(s1)), uint64(len(s2))
	AB := s1 + s2

	var wordLen = uint64(len(AB))
	x, y := l1, l2

	position := n

	var number string

	for {
		if position < y && position > x {
			position = position - x
			x, y = l1, l2
			continue
		}

		if position <= wordLen {
			number = string(AB[position-1])
			break
		}

		x, y = y, x+y
	}

	return number
}

func fibonacci(s1 string, s2 string) func() int {

	x, y := len(s1), len(s2)

	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	//s := "1415926535 8979323846 35"
	//s := "1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679 8214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196 104683731294243150"
	s := "123 4567 25"
	resS := strings.Split(s, " ")
	s1 := resS[0]
	s2 := resS[1]
	pos, _ := strconv.ParseUint(resS[2], 10, 64)

	res := fibonacciWords(s1, s2, pos)
	fmt.Println(res)
}
