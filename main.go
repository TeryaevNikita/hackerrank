package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func readFromFile() {
	//file, err := os.Open("C:/workspace/lara/go/input")
	//
	//if err != nil {
	//    log.Fatal(err)
	//}
	//
	//defer file.Close()
	//
	//var line []byte
	//var stringLine string
	//
	//reader := bufio.NewReader(file)
	//
	//line, _, err = reader.ReadLine()
	//stringLine = string(line)
	//scoresCount, err := strconv.Atoi(stringLine)
	//
	//var scores []int64
	//
	//stringLine, _ = reader.ReadString('\n')
	//scoresTemp := strings.Split(stringLine, " ")
	//
	//for i := 0; i < int(scoresCount); i++ {
	//    scoresTempStr := scoresTemp[i]
	//    if i == int(scoresCount)-1 {
	//        scoresTempStr = scoresTempStr[:len(scoresTempStr)-1]
	//    }
	//
	//    scoresItemTemp, err := strconv.ParseInt(scoresTempStr, 10, 64)
	//    checkError(err)
	//    scoresItem := int64(scoresItemTemp)
	//    scores = append(scores, scoresItem)
	//}
	//
	//line, _, err = reader.ReadLine()
	//stringLine = string(line)
	//aliceCount, err := strconv.Atoi(stringLine)
	//
	//var alice []int64
	//
	//stringLine, _ = reader.ReadString('\n')
	//aliceTemp := strings.Split(stringLine, " ")
	//
	//for i := 0; i < int(aliceCount); i++ {
	//    tempStr := aliceTemp[i]
	//    if i == int(aliceCount)-1 {
	//        tempStr = tempStr[:len(tempStr)-1]
	//    }
	//
	//    aliceItemTemp, err := strconv.ParseInt(tempStr, 10, 64)
	//    checkError(err)
	//    aliceItem := int64(aliceItemTemp)
	//    alice = append(alice, aliceItem)
	//}
	//
	////fmt.Println(scoresCount)
	////fmt.Println(aliceCount)
	////fmt.Println(len(scores))
	////fmt.Println(len(alice))
}

func plusMinus(arr []int32) {

	var nv, pv int
	var nf, pf, zf float32
	count := float32(len(arr))

	for _, value := range arr {
		switch {
		case value < 0:
			nv++
		case value > 0:
			pv++
		}
	}

	pf = float32(pv) / count
	nf = float32(nv) / count
	zf = 1.0 - nf - pf

	fmt.Printf("%.6f\n%.6f\n%.6f", pf, nf, zf)
}

func staircase(n int32) {
	var s string

	for i := n; i > 0; i-- {
		s = ""
		for j := int32(1); j <= n; j++ {
			if j < i {
				s += " "
			} else {
				s += "#"
			}
		}
		fmt.Println(s)
	}

}

func miniMaxSum(arr []int32) {
	var sum, min, max int64

	for i, value := range arr {
		value64 := int64(value)
		sum += value64

		if i == 0 {
			max = value64
			min = value64
		}

		if value64 > max {
			max = value64
		} else if value64 < min {
			min = value64
		}
	}

	fmt.Println(sum-max, sum-min)
}

func birthdayCakeCandles(ar []int32) int32 {

	var max, total int32

	for i, value := range ar {
		if i == 0 {
			max = value
		}

		if value > max {
			max = value
			total = 1
		} else if value == max {
			total++
		}
	}

	return total
}

func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */

	hh := s[:2]
	mm := s[3:5]
	ss := s[6:8]
	format := s[8:10]

	hour, _ := strconv.Atoi(hh)

	hour = hour % 12

	if format == "PM" {
		hour += 12
	}

	//hour = hour % 24
	hh = strconv.Itoa(hour)

	if hour < 10 {
		hh = "0" + hh
	}

	return fmt.Sprintf("%s:%s:%s", hh, mm, ss)
}

func boardCutting(cost_y []int32, cost_x []int32) int32 {

	var xSegments, ySegments uint64 = 1, 1

	sort.Slice(cost_y, func(i, j int) bool {
		return cost_y[i] > cost_y[j]
	})

	sort.Slice(cost_x, func(i, j int) bool {
		return cost_x[i] > cost_x[j]
	})

	yLen, xLen := len(cost_y), len(cost_x)
	length := yLen + xLen

	var result, yCost, xCost uint64

	for i, yPos, xPos := 0, 0, 0; i < length; i++ {
		if yPos < yLen {
			yCost = uint64(cost_y[yPos])
		} else {
			yCost = 0
		}

		if xPos < xLen {
			xCost = uint64(cost_x[xPos])
		} else {
			xCost = 0
		}

		if yCost >= xCost && yPos < yLen {
			result = (result + yCost*ySegments) % (1E9 + 7)
			yPos++
			xSegments++
		} else if xPos < xLen {
			result = (result + xCost*xSegments) % (1E9 + 7)
			xPos++
			ySegments++
		}

	}

	return int32(result % (1E9 + 7))
}

// NOT SOLVED
func reverseShuffleMerge(s string) string {

	lenS := len(s)
	//wordLen := lenS / 2

	reverseStr := make([]uint8, lenS)

	for i := range s {
		reverseStr[i] = s[lenS-i-1]
	}

	dict := make(map[uint8]int)
	//var keys []int32
	var startPos int
	var minChar = uint8('z')

	fmt.Println(reverseStr)
	fmt.Println(minChar)

	for i, value := range reverseStr {
		count, ok := dict[value]

		if !ok {
			count = 1
		} else {
			count++
		}

		if value < minChar {
			startPos = i
			minChar = value
		}

		dict[value] = count
	}

	//for char, value := range dict {
	//
	//    for j := startPos; j < lenS; j++ {
	//        if reverseStr[j] == char {
	//
	//        }
	//    }
	//}

	//for i, value := range s {
	//
	//}
	fmt.Println(startPos)
	fmt.Println(dict)

	// To create a map as input
	//m := make(map[int]string)
	//m[1] = "a"
	//m[2] = "c"
	//m[0] = "b"
	//
	//// To store the keys in slice in sorted order
	//var keys []int
	//for k := range m {
	//
	//    keys = append(keys, k)
	//}
	//fmt.Println(keys)
	//sort.Ints(keys)
	//fmt.Println(keys)
	//
	//// To perform the opertion you want
	//for _, k := range keys {
	//    fmt.Println("Key:", k, "Value:", m[k])
	//}

	return "1"
}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var roudGrade []int32
	var curGrade int32

	for _, value := range grades {
		if value >= 38 {
			left := value % 5
			if left >= 3 {
				curGrade = value - left + 5
			} else {
				curGrade = value
			}
		} else {
			curGrade = value
		}

		roudGrade = append(roudGrade, curGrade)
	}

	return roudGrade
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var fruitsAp, fruitsOr int32

	for _, ap := range apples {
		if a+ap >= s && a+ap <= t {
			fruitsAp++
		}
	}

	for _, or := range oranges {
		if b+or <= t && b+or >= s {
			fruitsOr++
		}
	}

	fmt.Println(fruitsAp)
	fmt.Println(fruitsOr)

}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	pos1, pos2 := x1, x2

	var dist, dist2 int32 = pos2 - pos1, 0

	if dist == 0 {
		return "YES"
	}

	var positive, newPositive bool

	if dist > 0 {
		positive = true
		newPositive = true
	}

	for {
		pos1 += v1
		pos2 += v2
		dist2 = pos2 - pos1

		if dist2 == 0 {
			return "YES"
		}

		if math.Abs(float64(dist2)) >= math.Abs(float64(dist)) {
			return "NO"
		}

		if dist2 > 0 {
			newPositive = true
		}

		if newPositive != positive {
			return "NO"
		}

	}

}

func _GCD(a, b int32) int32 {

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func _LCM(a, b int32) int32 {
	result := a * b / _GCD(a, b)
	return result
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var max, min, cnt int32

	for i, val := range a {
		if i == 0 {
			min = val
		}

		min = _LCM(min, val)
	}

	for i, val := range b {
		if i == 0 {
			max = val
			continue
		}

		max = _GCD(max, val)
	}

	for j := min; j <= max; j = j + min {
		if max%j == 0 {
			cnt++
		}
	}

	return cnt
}

func breakingRecords(scores []int32) []int32 {

	var arr1 []int32
	var min, max, countMin, countMax int32

	for i, value := range scores {
		if i == 0 {
			min = value
			max = value
			continue
		}

		if value < min {
			min = value
			countMin++
		}

		if value > max {
			max = value
			countMax++
		}
	}

	arr1 = append(arr1, countMax, countMin)

	return arr1
}

func hourglassSum(arr [][]int32) int32 {

	lenArr := len(arr)
	var sum, max int32

	for i := 0; i <= lenArr-3; i++ {
		for j := 0; j <= lenArr-3; j++ {
			sum = 0

			sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if i == 0 && j == 0 {
				max = sum
			}

			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func reverseArray(a []int32) []int32 {

	var length = len(a)
	for i := 0; i < length/2; i++ {
		a[i], a[length-i-1] = a[length-i-1], a[i]
	}

	return a
}

func matchingStrings(strings []string, queries []string) []int32 {

	var res []int32
	counts := make(map[string]int)

	for _, value := range strings {
		counts[value]++
	}

	for _, value := range queries {
		count, ok := counts[value]

		if !ok {
			count = 0
		}

		res = append(res, int32(count))

	}

	return res
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	var lastAnswer int32 = 0
	var lastAnswers []int32

	s := make([][]int32, n)

	for _, query := range queries {
		q, x, y := query[0], query[1], query[2]

		seq := (x ^ lastAnswer) % n

		fmt.Println("seq=", seq)

		if q == 1 {
			s[seq] = append(s[seq], y)
		} else {
			seqLen := int32(len(s[seq]))
			seqIndex := y % seqLen
			lastAnswer = s[seq][seqIndex]
			lastAnswers = append(lastAnswers, lastAnswer)
		}
	}

	return lastAnswers
}

func arrayLeftRotation(n int32, arr []int32) []int32 {

	arrLen := len(arr)
	result := make([]int32, arrLen)

	for i, value := range arr {
		index := (int32(i) - n + int32(arrLen)) % int32(arrLen)
		result[index] = value
	}

	return result
}

func arrayManipulation(n int32, queries [][]int32) int64 {

	var max, current int64

	var arr = make([]int64, n)

	for _, query := range queries {
		a, b, k := int64(query[0]), int64(query[1]), int64(query[2])

		arr[a-1] += k
		if b < int64(n) {
			arr[b] -= k
		}
	}

	fmt.Println(arr)
	for _, value := range arr {
		current += value

		if current > max {
			max = current
		}
	}

	return max
}

func birthday(s []int32, d int32, m int32) int32 {

	var count int32

	for i := int32(0); i < int32(len(s))-m+1; i++ {
		sum := int32(0)

		for j := int32(0); j < m; j++ {
			fmt.Println(i + j)

			sum += s[i+j]
		}

		if sum == d {
			count++
		}

	}
	return count
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	arrLen := int32(len(ar))

	for i := int32(0); i < arrLen-1; i++ {
		value1 := ar[i]
		for j := i + 1; j < arrLen; j++ {
			if (value1+ar[j])%k == 0 {
				count++
			}
		}
	}

	return count
}

func formingMagicSquare(s [][]int32) int32 {

	var existingMagic [][]int32

	magicTemplate := [3][3]int32{
		{1, -4, 3},
		{2, 0, -2},
		{-3, 4, -1},
	}

	magicOrigin := func(arr [3][3]int32) [3][3]int32 {
		for i, value := range arr {
			for j := range value {
				arr[i][j] += 5
			}
		}
		return arr
	}(magicTemplate)

	rotate := func(arr [3][3]int32) [3][3]int32 {
		var res [3][3]int32

		for i, value := range arr {
			for j := range value {
				res[j][2-i] = arr[i][j]
			}
		}
		return res
	}

	mirror := func(arr [3][3]int32) [3][3]int32 {
		var res [3][3]int32

		for i, value := range arr {
			for j := range value {
				res[i][2-j] = arr[i][j]
			}
		}
		return res
	}

	flat := func(arr [3][3]int32) []int32 {
		var res []int32
		for i, value := range arr {
			for j := range value {
				res = append(res, arr[i][j])
			}
		}
		return res
	}

	var curMagic [3][3]int32

	curMagic = magicOrigin
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = mirror(magicOrigin)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	curMagic = rotate(curMagic)
	existingMagic = append(existingMagic, flat(curMagic))

	var sum, min int32
	for ii, origins := range existingMagic {
		sum = 0
		for i, row := range s {
			for j, value := range row {
				dif := value - origins[3*i+j]
				if dif < 0 {
					dif = -dif
				}
				sum += dif
			}
		}

		if ii == 0 {
			min = sum
		}

		if sum < min {
			min = sum
		}
	}

	return min
}

func pickingNumbers(a []int32) int32 {
	// Write your code here
	arrData := make(map[int32]int32)
	var keys []int

	for _, value := range a {
		arrData[value]++
	}

	for i := range arrData {
		keys = append(keys, int(i))
	}

	sort.Ints(keys)

	var lastKey, max, count, lastValue int32

	for i, key := range keys {

		curKey := int32(key)
		curValue := arrData[curKey]

		if i == 0 {
			lastKey = curKey
			lastValue = curValue
			max = curValue
			continue
		}

		count = curValue

		if curKey-lastKey == 1 {
			count += lastValue
		}

		if count > max {
			max = count
		}
		lastKey = curKey
		lastValue = curValue
	}

	return max
}

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {

	tip := meal_cost * float64(tip_percent) / 100.0
	tax := meal_cost * float64(tax_percent) / 100.0

	totalCost := meal_cost + tip + tax

	fmt.Println(math.Round(totalCost))
}

func migratoryBirds(arr []int32) int32 {
	arrData := make(map[int32]int32)
	var keys []int

	for _, value := range arr {
		arrData[value]++
	}

	for i := range arrData {
		keys = append(keys, int(i))
	}

	sort.Ints(keys)

	var maxKey, max int32

	for i, key := range keys {

		curKey := int32(key)
		curValue := arrData[curKey]

		if i == 0 {
			maxKey = curKey
			max = curValue
			continue
		}

		if curValue > max {
			max = curValue
			maxKey = curKey
		}

	}

	return maxKey
}

func dayOfProgrammer(year int32) string {

	var dayX, month, day int32 = 256, 0, 0

	monthDays := [12]int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year == 1918 {
		dayX += 13
	}

	for i, value := range monthDays {
		if dayX < value {
			month = int32(i) + 1
			day = dayX
			break
		}
		dayX -= value

		if i == 1 {
			if year%400 == 0 {
				dayX--
			} else if year%4 == 0 && (year%100 != 0 || year < 1918) {
				dayX--
			}
		}
	}

	return fmt.Sprintf("%d.%02d.%d", day, month, year)
}

func climbingLeaderboard_(scores []int64, alice []int64) []int64 {

	type Point struct {
		Point, Pos int64
	}

	var points []Point

	var position int64 = 1

	var prev int64
	for i, value := range scores {
		if i == 0 {
			prev = value
		}

		if prev == value {
			continue
		}

		points = append(points, Point{prev, position})
		prev = value
		position++
	}

	points = append(points, Point{prev, position})

	findPosition := func(current int64, arr []Point) int64 {

		var start, end int64 = 0, int64(len(points) - 1)
		var middle, position int64
		var currentPoint Point

		for {

			if end == start {
				currentPoint = arr[end]

				if currentPoint.Point < current || currentPoint.Point == current {
					position = currentPoint.Pos
				} else {
					position = currentPoint.Pos + 1
				}
				break
			}

			middle = (start + end) / 2

			currentPoint = arr[middle]

			if currentPoint.Point == current {
				position = currentPoint.Pos
				break
			} else if currentPoint.Point > current {
				start = middle + 1
			} else {
				end = middle - 1
				if end < start {
					end = start
				}
			}
		}

		return position
	}

	var result []int64
	for _, value := range alice {
		result = append(result, findPosition(value, points))
	}

	return result
}

func climbingLeaderboard(scores []int64, alice []int64) []int64 {

	positions := make(map[int64]int64)

	var position int64 = 1

	for _, value := range scores {
		if _, ok := positions[value]; !ok {
			positions[value] = position
			position++
		}
	}

	findPosition := func(current int64, baseStart int64, arr []int64) int64 {

		var start, end = baseStart, int64(len(arr) - 1)
		var middle, position int64
		var currentValue int64

		for {
			if end == start {
				currentValue = arr[end]

				if currentValue < current || currentValue == current {
					position = positions[currentValue]
				} else {
					position = positions[currentValue] + 1
				}
				break
			}

			middle = (start + end) / 2

			currentValue = arr[middle]

			if currentValue == current {
				position = positions[currentValue]
				break
			} else if currentValue > current {
				start = middle + 1
			} else {
				end = middle - 1
				if end < start {
					end = start
				}
			}
		}

		return position
	}

	var result []int64
	for _, value := range alice {
		result = append(result, findPosition(value, 0, scores))
	}

	return result
}

func bonAppetit(bill []int32, k int32, b int32) {

	var total, anna int32
	for i, price := range bill {
		total += price
		if int32(i) == k {
			anna = price
		}
	}

	delta := b - (total-anna)/2

	if delta != 0 {
		fmt.Println(delta)
	} else {
		fmt.Println("Bon Appetit")
	}
}

func sockMerchant(n int32, ar []int32) int32 {

	var count int32

	arrMap := make(map[int32]int32)

	for _, color := range ar {
		arrMap[color]++

		if arrMap[color]%2 == 0 {
			arrMap[color] = 0
			count++
		}
	}

	return count
}

func pageCount(n int32, p int32) int32 {

	var result int32
	totalCounts := n/2 + 1

	count := p/2 + 1

	fromRight, fromLeft := count-1, totalCounts-count

	if fromRight < fromLeft {
		result = fromRight
	} else {
		result = fromLeft
	}

	return result
}

func findDigits(n int32) int32 {

	var d, res int32
	m := n
	for {
		if m == 0 {
			break
		}

		d = m % 10
		m = m / 10

		if d == 0 {
			continue
		}
		if n%d == 0 {
			res++
		}
	}

	return res
}

func extraLongFactorials(n int32) {

	var result string

	longSum := func(a map[int]int, b map[int]int) map[int]int {
		la := len(a)
		lb := len(b)
		var maxLen int
		result := make(map[int]int)

		if la > lb {
			maxLen = la
		} else {
			maxLen = lb
		}

		var valA, valB, left int
		var ok bool
		for i := 0; i < maxLen; i++ {

			if valA, ok = a[i]; !ok {
				valA = 0
			}

			if valB, ok = b[i]; !ok {
				valB = 0
			}

			sum := valA + valB + left

			number := sum % 10

			left = sum / 10
			result[i] = number
		}

		if left > 0 {
			result[maxLen] = left
		}

		return result
	}

	multiple := func(a string, b string) string {

		var mult int
		var res string

		multMap := make(map[int]int)
		resMap := make(map[int]int)

		for i := len(b) - 1; i >= 0; i-- {

			for j := len(a) - 1; j >= 0; j-- {
				m1, _ := strconv.Atoi(string(a[j]))
				m2, _ := strconv.Atoi(string(b[i]))

				mult = m1 * m2

				pos := 0

				multMap = make(map[int]int)

				for jj := 0; jj < (len(b)+len(a)-2)-i-j; jj++ {
					multMap[pos] = 0
					pos++
				}

				for {
					dec := mult % 10
					mult = mult / 10
					multMap[pos] = dec
					if mult == 0 {
						break
					}
					pos++
				}

				resMap = longSum(resMap, multMap)

			}
		}

		for ii := len(resMap) - 1; ii >= 0; ii-- {
			res += fmt.Sprintf("%d", resMap[ii])
		}

		return res
	}

	for i := 1; i <= int(n); i++ {
		if i == 1 {
			result = "1"
		} else {
			result = multiple(result, strconv.Itoa(i))
		}
	}

	fmt.Println(result)
}

func extraLongFactorialsLib(n int32) {
	factorial := new(big.Int)

	factorial.MulRange(1, int64(n))

	fmt.Println(factorial)
}

func countingValleys(n int32, s string) int32 {

	high, count := 0, 0

	for _, value := range s {

		if value == 'U' {
			high++
		} else {
			high--
		}

		if high == 0 && value == 'U' {
			count++
		}
	}

	return int32(count)
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {

	sort.Slice(keyboards, func(i, j int) bool {
		return keyboards[j] > keyboards[i]
	})

	sort.Slice(drives, func(i, j int) bool {
		return drives[j] > drives[i]
	})

	var min int32 = -1

	for i, priceK := range keyboards {
		for j, priceD := range drives {

			if i == 0 && j == 0 && priceK+priceD > b {
				return -1
			}

			if priceK+priceD > b {
				continue
			}

			if priceK+priceD > min {
				min = priceK + priceD
			}
		}
	}

	return min
}

func catAndMouse(x int32, y int32, z int32) string {
	var res string

	cat1 := x - z
	if cat1 < 0 {
		cat1 = -cat1
	}
	cat2 := y - z

	if cat2 < 0 {
		cat2 = -cat2
	}

	switch {
	case cat1 < cat2:
		res = "Cat A"
	case cat2 < cat1:
		res = "Cat B"
	case cat1 == cat2:
		res = "Mouse C"

	}
	return res
}

func hurdleRace(k int32, height []int32) int32 {
	var res, max int32 = 0, -1

	for _, value := range height {
		if value > max {
			max = value
		}
	}

	if max > k {
		res = max - k
	}

	return res
}

func designerPdfViewer(h []int32, word string) int32 {
	var max int32
	for _, letter := range word {
		if h[letter-97] > max {
			max = h[letter-97]
		}
	}

	return max * int32(len(word))
}

func utopianTree(n int32) int32 {

	grow := func() func() int32 {
		var h, period int32 = 1, 0
		return func() int32 {
			if period != 0 {
				if period%2 == 0 {
					h += 1
				} else {
					h *= 2
				}
			}
			period++
			return h
		}
	}()

	var res int32

	for i := int32(0); i <= n; i++ {
		res = grow()
	}

	return res
}

func angryProfessor(k int32, a []int32) string {
	var count int32

	for _, value := range a {
		if value <= 0 {
			count++
			if count >= k {
				break
			}
		}
	}

	res := "YES"

	if count >= k {
		res = "NO"
	}

	return res
}

func encryption(s string) string {
	length := len(s)

	resLen := math.Sqrt(float64(length))
	rowLen, columnLen := int(math.Floor(resLen)), int(math.Ceil(resLen))

	if rowLen*columnLen < length {
		rowLen++
	}

	array := make([][]byte, rowLen)
	for i := range array {
		array[i] = make([]byte, columnLen)
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			position := i*columnLen + j
			if position >= length {
				continue
			}

			array[i][j] = s[i*columnLen+j]
		}
	}
	fmt.Println(array)

	res := ""
	for j := 0; j < columnLen; j++ {
		for i := 0; i < rowLen; i++ {
			value := array[i][j]
			if value == 0 {
				continue
			}
			res += fmt.Sprint(string(array[i][j]))
		}
		res += " "
	}

	return res
}

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here

	remainder := make([]int32, k)

	for _, value := range s {
		rem := int(value % k)
		remainder[rem]++
	}

	var count int32

	for i := int32(1); i < (k+1)/2; i++ {
		if remainder[i] > remainder[k-i] {
			count += remainder[i]
		} else {
			count += remainder[k-i]
		}
	}

	if zeroRem := remainder[0]; zeroRem > 0 {
		count++
	}

	if k%2 == 0 {
		if halfRem := remainder[k/2]; halfRem > 0 {
			count++
		}
	}

	return count
}

func beautifulDays(i int32, j int32, k int32) int32 {

	var count int32

	reverse := func(n int32) (m int32) {

		for n > 0 {
			reminder := n % 10
			m *= 10
			n /= 10
			m += reminder
		}
		return
	}

	for ll := i; ll <= j; ll++ {
		dif := ll - reverse(ll)
		if dif < 0 {
			dif = -dif
		}

		if dif%k == 0 {
			count++
		}
	}

	return count
}

func viralAdvertising(n int32) int32 {

	media := func() func() int32 {
		var sum, shared, liked int32 = 0, 5, 0
		return func() int32 {
			liked = shared / 2
			shared = liked * 3
			sum += liked
			return sum
		}
	}()

	var res int32
	for i := int32(1); i <= n; i++ {
		res = media()
	}

	return res
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return (s+m-2)%n + 1
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	length := int32(len(a))

	b := make([]int32, length)

	for i := range a {
		b[(int32(i)+k)%length] = a[i]
	}
	fmt.Println(a, b)

	var c []int32

	for _, value := range queries {
		c = append(c, b[value])
	}

	return c
}

func jumpingOnClouds(c []int32, k int32) int32 {

	var points, pos, length int32 = 100, 0, int32(len(c))

	for {
		pos = (pos + k) % length
		points--

		if c[pos] == 1 {
			points -= 2
		}

		if pos == 0 {
			break
		}
	}

	return points
}

func appendAndDelete(s string, t string, k int32) string {

	sLen, tLen := len(s), len(t)

	var sameLen, removeItems, appendItem int32 = int32(sLen), 0, 0

	for i, value := range s {
		if i >= tLen || value != int32(t[i]) {
			sameLen = int32(i)
			break
		}
	}

	removeItems = int32(sLen) - sameLen
	appendItem = int32(tLen) - sameLen

	left := k - (removeItems + appendItem)
	if left < 0 {
		return "No"
	}

	if left%2 == 1 && left < sameLen*2 {
		return "No"
	}

	return "Yes"
}

func stones(n int32, a int32, b int32) []int32 {
	var sum int32
	var arr []int32

	values := make(map[int32]bool)
	if a > b {
		a, b = b, a
	}

	m := int(n - 1)
	fmt.Println(m)

	for i := 0; i <= m; i++ {
		sum = int32(i)*b + (int32(m-i))*a

		if _, ok := values[sum]; !ok {
			values[sum] = true
			arr = append(arr, sum)
		}
	}

	return arr
}

func cavityMap(grid []string) []string {

	length := len(grid)

	var result []string
	var str string

	result = make([]string, length)

	for i, row := range grid {
		str = ""
		for j := range row {

			if j == 0 || j == length-1 || i == 0 || i == length-1 {
				str += string(grid[i][j])
				continue
			}
			currentDeep := grid[i][j]

			left, top, right, bot := grid[i][j-1], grid[i-1][j], grid[i][j+1], grid[i+1][j]

			if currentDeep <= left || currentDeep <= top || currentDeep <= right || currentDeep <= bot {
				str += string(grid[i][j])
				continue
			}
			str += "X"
		}
		result[i] = str
	}

	return result
}

func squares(a int32, b int32) int32 {
	left, right := int32(math.Floor(math.Sqrt(float64(a)))), int32(math.Floor(math.Sqrt(float64(b))))

	count := right - left

	if left*left == a {
		count++
	}

	return count
}

func dictionariesAndMaps(book map[string]string, names []string) {

	for _, value := range names {
		if phone, ok := book[value]; ok {
			fmt.Printf("%s=%s\n", value, phone)
		} else {
			fmt.Println("Not found")
		}
	}
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {

	dDif, mDif, yDif := d2-d1, m2-m1, y2-y1

	if yDif < 0 {
		return 10000
	} else if yDif > 0 {
		return 0
	}

	if mDif < 0 {
		return -500 * mDif
	} else if mDif > 0 {
		return 0
	}

	if dDif < 0 {
		return -15 * dDif
	} else if dDif > 0 {
		return 0
	}

	return 0
}

func cutTheSticks(arr []int32) []int32 {

	store := make(map[int32]int32)
	var storeKey []int32

	for _, value := range arr {
		store[value]++
		if store[value] == 1 {
			storeKey = append(storeKey, value)
		}
	}

	sort.Slice(storeKey, func(i, j int) bool {
		return storeKey[int32(i)] < storeKey[int32(j)]
	})

	var result []int32
	count := int32(len(arr))
	for _, value := range storeKey {
		result = append(result, count)
		count -= store[value]
	}

	return []int32{}
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	var ox, oy int32

	var top, bot, left, right, rightTop, leftBot, leftTop, rightBot [2]int32
	qx, qy := r_q, c_q

	mainDiag, secondDiag := qx-qy, qx+qy

	top = [2]int32{n + 1, qy}
	bot = [2]int32{0, qy}
	left = [2]int32{qx, 0}
	right = [2]int32{qx, n + 1}

	if mainDiag >= 0 {
		rightTop = [2]int32{n + 1, n + 1 - mainDiag}
		leftBot = [2]int32{mainDiag, 0}
	} else {
		rightTop = [2]int32{n + 1 + mainDiag, n + 1}
		leftBot = [2]int32{0, -mainDiag}
	}

	if secondDiag >= n+1 {
		leftTop = [2]int32{n + 1, secondDiag - (n + 1)}
		rightBot = [2]int32{secondDiag - (n + 1), n + 1}
	} else {
		leftTop = [2]int32{secondDiag, 0}
		rightBot = [2]int32{0, secondDiag}
	}

	for _, obstacle := range obstacles {
		ox, oy = obstacle[0], obstacle[1]

		// row
		if ox == qx {
			if oy > qy {
				if oy < right[1] {
					right = [2]int32{ox, oy}
				}
			} else if oy < qy {
				if oy > left[1] {
					left = [2]int32{ox, oy}
				}
			}
			continue
		}

		//column
		if oy == qy {
			if ox > qx {
				if ox < top[0] {
					top = [2]int32{ox, oy}
				}
			} else if ox < qx {
				if ox > bot[0] {
					bot = [2]int32{ox, oy}
				}
			}
			continue
		}

		// main diagonal
		if ox-oy == mainDiag {
			if ox > qx {
				if rightTop[0] == 0 || ox < rightTop[0] {
					rightTop = [2]int32{ox, oy}
				}
			} else if ox < qx {
				if leftBot[0] == 0 || ox > leftBot[0] {
					leftBot = [2]int32{ox, oy}
				}
			}
		}

		// second diagonal
		if ox+oy == secondDiag {
			if ox < qx {
				if rightBot[0] == 0 || ox > rightBot[0] {
					rightBot = [2]int32{ox, oy}
				}
			} else if ox > qx {
				if leftTop[0] == 0 || ox < leftTop[0] {
					leftTop = [2]int32{ox, oy}
				}
			}
		}
	}

	var squares int32

	toTop := top[0] - qx - 1
	toBot := qx - bot[0] - 1
	toRight := right[1] - qy - 1
	toLeft := qy - left[1] - 1

	toRightTop := rightTop[0] - qx - 1
	toLeftBot := qx - leftBot[0] - 1
	toRightBot := rightBot[1] - qy - 1
	toLeftTop := qy - leftTop[1] - 1

	squares = toTop + toBot + toRight + toLeft + toRightTop + toLeftBot + toRightBot + toLeftTop

	return squares
}

func chocolateFeast(n int32, c int32, m int32) int32 {
	wrappers := n / c
	var count, wCount int32 = wrappers, 0

	for wrappers >= m {
		wCount = wrappers / m
		wrappers = wCount + (wrappers - wCount*m)
		count += wCount
	}

	return count
}

func matrixRotation(matrix [][]int32, r int32) {

	var result [][]int32

	n := int32(len(matrix))
	m := int32(len(matrix[0]))
	minWidth := n

	if m < minWidth {
		minWidth = m
	}

	result = make([][]int32, n)
	for i := int32(0); i < n; i++ {
		result[i] = make([]int32, m)
	}

	var width, height, newI, newJ int32

	for circleLevel := int32(0); circleLevel < minWidth/2; circleLevel++ {
		width, height = m-2*circleLevel, n-2*circleLevel
		loopSize := 2*width + 2*height - 4
		step := r % loopSize

		circleFlat := make([][]int32, loopSize)

		top, left, bot, right := circleLevel, circleLevel, n-circleLevel, m-circleLevel
		i, j := top, left
		direct := "down"

		for k := int32(0); k < loopSize; k++ {
			circleFlat[k] = []int32{i, j}
			if direct == "down" {
				i++
				if i >= bot {
					i--
					j++
					direct = "right"
				}
				continue
			}

			if direct == "right" {
				j++
				if j >= right {
					j--
					i--
					direct = "up"
				}
				continue
			}

			if direct == "up" {
				i--
				if i < top {
					i++
					j--
					direct = "left"
				}
				continue
			}

			if direct == "left" {
				j--
				if j < left {
					j++
					i--
					direct = "down"
				}
				continue
			}

		}

		for ind, point := range circleFlat {
			i, j := point[0], point[1]

			newPoint := circleFlat[(int32(ind)+step)%loopSize]
			newI, newJ = newPoint[0], newPoint[1]

			result[newI][newJ] = matrix[i][j]
		}
	}

	for _, row := range result {
		for _, value := range row {
			fmt.Print(value)
			fmt.Print(" ")

		}
		fmt.Println()
	}

}

func repeatedString(s string, n int64) int64 {

	length := int64(len(s))
	countA := int64(strings.Count(s, "a"))

	wordCount := n / length
	left := n % length

	s1 := s[:left]

	count := countA*wordCount + int64(strings.Count(s1, "a"))

	return count
}

func organizingContainers(container [][]int32) string {

	n := len(container)
	result := "Possible"
	var countCol, countRow int32
	var cols, rows []int32

	for i := 0; i < n; i++ {
		countCol = 0
		countRow = 0
		for j := 0; j < n; j++ {
			countCol += container[j][i]
			countRow += container[i][j]
		}

		cols = append(cols, countCol)
		rows = append(rows, countRow)
	}

	sort.Slice(cols, func(i, j int) bool {
		return cols[int32(i)] > cols[int32(j)]
	})

	sort.Slice(rows, func(i, j int) bool {
		return rows[int32(i)] > rows[int32(j)]
	})

	for i, value := range cols {
		if value != rows[i] {
			result = "Impossible"
			break
		}
	}

	return result
}

func biggerIsGreater(w string) string {

	length := len(w)
	var pos, newPos = -1, 0
	var left []byte
	var result string
	var newValue byte = 'z' + 1

	for i := length - 1; i > 0; i-- {
		if pos < 0 {
			left = append(left, w[i])
		}

		if w[i] > w[i-1] && pos < 0 {
			pos = i - 1
		}
	}

	if pos == -1 {
		return "no answer"
	}

	curValue := w[pos]

	for i, value := range left {
		if value > curValue && value < newValue {
			newValue = value
			newPos = i
		}
	}

	curValue, left[newPos] = left[newPos], curValue

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	for i := 0; i < pos; i++ {
		result += string(w[i])
	}

	result += string(curValue)

	for _, value := range left {
		result += string(value)
	}

	return result
}

func jumpingOnClouds2(c []int32) int32 {

	var pos, jumps int32
	length := int32(len(c))

	for {
		if pos+2 < length && c[pos+2] != 1 {
			jumps++
			pos += 2
			continue
		}

		if pos+1 < length && c[pos+1] != 1 {
			jumps++
			pos += 1
			continue
		}

		if pos >= length-1 {
			break
		}

	}

	return jumps
}

func equalizeArray(arr []int32) int32 {
	length := len(arr)
	itemCount := make(map[int32]int32)

	for _, val := range arr {
		itemCount[val]++
	}

	var max int32

	for _, val := range itemCount {
		if val > max {
			max = val
		}
	}

	return int32(length) - max
}

func acmTeam(topic []string) []int32 {

	var MaxSize = 64

	countStrings := len(topic)
	topics := int(math.Ceil(float64(len(topic[0])) / float64(MaxSize)))

	intArr := make([][]uint64, countStrings)

	var rowArr []uint64
	var start, end int

	for i, stringRow := range topic {
		sLen := len(stringRow)
		start = 0
		end = sLen
		if end > MaxSize {
			end = MaxSize
		}

		for {
			value, _ := strconv.ParseUint(stringRow[start:end], 2, 64)
			rowArr = append(rowArr, value)

			start = end
			end += MaxSize

			if end > sLen {
				end = sLen
			}

			if start >= sLen {
				break
			}
		}

		intArr[i] = rowArr
		rowArr = []uint64{}
	}

	var maxOnes, countMax, curCount int32

	for i := 0; i < countStrings-1; i++ {
		for j := i + 1; j < countStrings; j++ {
			curCount = 0
			for ii := 0; ii < topics; ii++ {
				value := intArr[i][ii] | intArr[j][ii]
				curCount += int32(bits.OnesCount64(value))
			}

			if curCount > countMax {
				countMax = curCount
				maxOnes = 0
			}

			if curCount == countMax {
				//fmt.Println(curCount, i+1,j+1)

				maxOnes++
			}
		}
	}

	return []int32{countMax, maxOnes}
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here

	b64, w64, bc64, wc64, z64 := int64(b), int64(w), int64(bc), int64(wc), int64(z)
	bPrice, wPrice := bc64, wc64

	if wc64+z64 < bPrice {
		return wc64*(b64+w64) + z64*b64
	}

	if bc64+z64 < wPrice {
		return bc64*(b64+w64) + z64*w64
	}

	return b64*bc64 + w64*wc64
}

func kaprekarNumbers(p int32, q int32) {
	var d int
	p64, q64 := int64(p), int64(q)

	var invalid = true

	for i := p64; i <= q64; i++ {
		d = int(math.Log10(float64(i))) + 1

		n := i * i
		rSlug := n % int64(math.Pow10(d))
		lSlug := n / int64(math.Pow10(d))

		if i == rSlug+lSlug {
			invalid = false
			fmt.Print(i)
			fmt.Print(" ")
		}
	}

	if invalid {
		fmt.Print("INVALID RANGE")
	}
}

func beautifulTriplets(d int32, arr []int32) int32 {

	var loops int32 = 2
	var maxPos, step, pos int32
	length := int32(len(arr))
	var value, candidate, counts int32

	for i := int32(0); i < length-loops; i++ {
		value = arr[i]

		j := i + 1
		step = 0

		for {
			maxPos = length - loops + step + 1
			find := false

			for l := j; l < maxPos; l++ {
				candidate = arr[l]
				if value+d == candidate {
					find = true
					pos = l
					break
				}

				if candidate > value+d {
					break
				}

			}

			if find {
				value = candidate
				step++
				j = pos
				if step == loops {
					counts++
					break
				}
			} else {
				break
			}
		}
	}

	return counts
}

func minimumDistances(a []int32) int32 {
	arrMap := make(map[int32][]int)

	for i, value := range a {
		arrMap[value] = append(arrMap[value], i)
	}

	min := len(a) + 1

	for _, data := range arrMap {
		if len(data) < 2 {
			continue
		}

		for i := 0; i < len(data)-1; i++ {
			res := data[i+1] - data[i]

			if res < min {
				min = res
			}
		}
	}

	if min > len(a) {
		min = -1
	}

	return int32(min)
}

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var sum, counts int32

	price := p
	for {

		if sum+price <= s {
			counts++
			sum += price
		} else {
			break
		}

		price -= d
		if price < m {
			price = m
		}
	}

	return counts
}

func timeInWords(h int32, m int32) string {
	var timeMap = map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21: "twenty one",
		22: "twenty two",
		23: "twenty three",
		24: "twenty four",
		25: "twenty five",
		26: "twenty six",
		27: "twenty seven",
		28: "twenty eight",
		29: "twenty nine",
		30: "half",
	}

	var result string
	var mInd, hInd int32
	var past bool

	if m == 0 {
		result = timeMap[h] + " o' clock"
	} else {

		if m > 30 {
			mInd = 60 - m
			hInd = h + 1
			past = false
		} else {
			mInd = m
			hInd = h
			past = true
		}
		result = timeMap[mInd]

		if m%15 != 0 {
			if m == 1 {
				result += " minute"
			} else {
				result += " minutes"
			}
		}

		if past {
			result += " past"
		} else {
			result += " to"
		}

		result += " " + timeMap[hInd]
	}

	return result
}

func almostSorted(arr []int32) {

	var left, right int
	isOk := true
	length := len(arr)

	for i := 0; i < length-1; i++ {
		if arr[i+1] < arr[i] {
			isOk = false
			left = i
			break
		}
	}

	//
	if isOk {
		fmt.Println("yes")
		return
	}

	for i := length - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			right = i
			break
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	checkArr := true
	for i := 0; i < length-1; i++ {
		if arr[i+1] < arr[i] {
			checkArr = false
			break
		}
	}

	if checkArr {
		fmt.Println("yes")
		fmt.Println("swap", left+1, right+1)
		return
	}

	arr[left], arr[right] = arr[right], arr[left]

	for i := 0; i <= (right-left)/2; i++ {
		arr[left+i], arr[right-i] = arr[right-i], arr[left+i]
	}

	checkArr = true
	for i := 0; i < length-1; i++ {
		if arr[i+1] < arr[i] {
			checkArr = false
			break
		}
	}

	if checkArr {
		fmt.Println("yes")
		fmt.Println("reverse", left+1, right+1)
		return
	}

	fmt.Println("no")
}

func workbook(n int32, k int32, arr []int32) int32 {
	var taskPages, page, maxTask int32 = 0, 1, 0

	book := make(map[int32][2]int32)

	for _, value := range arr {

		left := value % k
		taskPages = value / k
		if left != 0 {
			taskPages++
		}

		for i := int32(0); i < taskPages; i++ {
			maxTask = (i + 1) * k
			if maxTask > value {
				maxTask = value
			}
			book[page] = [2]int32{1 + i*k, maxTask}
			page++
		}
	}

	var count int32
	for i, value := range book {
		min, max := value[0], value[1]

		if i >= min && i <= max {
			count++
		}
	}

	return count
}

func insertionSort1(n int32, arr []int32) {

	value, pos := arr[n-1], n-1
	var insert = false
	for {
		candidatePos := pos - 1

		if candidatePos >= 0 && arr[candidatePos] > value {
			arr[pos] = arr[candidatePos]
			pos--
		} else {
			arr[pos] = value
			insert = true
		}

		for _, arrV := range arr {
			fmt.Print(arrV)
			fmt.Print(" ")
		}
		fmt.Println()

		if insert {
			break
		}

	}

}

func insertionSort2(n int32, arr []int32) {

	_insertionSort := func(n int32, arr []int32) {

		value := arr[n]
		for {
			if n > 0 && value < arr[n-1] {
				arr[n] = arr[n-1]
				n--
			} else {
				arr[n] = value
				break
			}
		}

	}

	for i := int32(1); i < n; i++ {
		_insertionSort(i, arr)

		for _, arrV := range arr {
			fmt.Print(arrV)
			fmt.Print(" ")
		}
		fmt.Println()

	}

}

func fairRations(B []int32) int32 {

	length := int32(len(B))

	var counts int32

	for i := int32(0); i < length-1; i++ {
		if B[i]%2 == 0 {
			continue
		}

		counts += 2
		B[i]++
		B[i+1]++
	}

	if B[length-1]%2 != 0 {
		return -1
	}

	return counts

}

func gridSearch(G []string, P []string) string {

	rowsG, rowsP := int32(len(G)), int32(len(P))
	colsG, colsP := int32(len(G[0])), int32(len(P[0]))

	leftTop := P[0][0]

	checkPattern := func(G []string, P []string) func(row, col int32) bool {
		rowP, colP := int32(len(P)), int32(len(P[0]))

		return func(row, col int32) bool {

			for i := int32(0); i < rowP; i++ {
				for j := int32(0); j < colP; j++ {
					if P[i][j] != G[i+row][j+col] {
						return false
					}
				}
			}

			return true
		}
	}(G, P)

	for i := int32(0); i <= rowsG-rowsP; i++ {
		for j := int32(0); j <= colsG-colsP; j++ {
			if G[i][j] == leftTop {
				res := checkPattern(i, j)

				if res {
					return "YES"
				}
			}
		}
	}

	return "NO"
}

func bigSorting(unsorted []string) []string {

	var valueI, valueJ string
	var lengthI, lengthJ int

	sort.Slice(unsorted, func(i, j int) bool {

		valueI, valueJ = unsorted[i], unsorted[j]

		lengthI = len(valueI)
		lengthJ = len(valueJ)

		if lengthI != lengthJ {
			return lengthI < lengthJ
		}

		for l := 0; l < lengthI; l++ {
			if valueI[l] != valueJ[l] {
				return valueI[l] < valueJ[l]
			}
		}

		return false
	})

	return unsorted
}

func introTutorial(V int32, arr []int32) int32 {
	findPos := func(arr []int32, left, right, value int32) int32 {

		for {
			pos := (right + left) / 2

			if value == arr[pos] {
				return pos
			}

			if value < arr[pos] {
				right = pos - 1
			} else {
				left = pos + 1
			}
		}

	}

	return findPos(arr, 0, int32(len(arr))-1, V)
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//
	//tempString := readLine(reader)
	//n, _ := strconv.ParseInt(tempString, 10, 32)
	//
	//book := make(map[string]string)
	//for i:=int64(0); i < n; i++ {
	//    tempString = readLine(reader)
	//    tempArr := strings.Split(tempString, " ")
	//    book[tempArr[0]] = tempArr[1]
	//}
	//
	//names := make([]string, n)
	//for i:=0; i < int(n); i++ {
	//    tempString = readLine(reader)
	//    names = append(names, tempString)
	//}

	var arr [][]int32
	arr = append(arr,
		[]int32{2, 6, 8},
		[]int32{3, 5, 7},
		[]int32{1, 8, 1},
		[]int32{5, 9, 15},
	)
	//arr = append(arr, 10, 6, 15, 20, 30, 5, 7)
	//arr = append(arr, 1, 4, 5, 7)
	//arr = append(arr, 1, 4, 5)
	//arr = append(arr, 1, 4)
	//arr = append(arr,
	//	"1",
	//	"22",
	//	"333",
	//	"4444",
	//	"55555",
	//	"666666",
	//)

	//arr1 = append(arr1,
	//    //[]int32{1, 3, 1},
	//    //[]int32{2, 1, 2},
	//    //[]int32{3, 3, 3},
	//    []int32{0, 2, 1},
	//    []int32{1, 1, 1},
	//    []int32{2, 0, 0},
	//    //[]int32{9, 10, 11, 12, 1212},
	//    //[]int32{13, 14, 15, 16, 1616},
	//    //[]int32{10, 9, 8, 7},
	//)

	//insertionSort(arr)
	result := arrayManipulation(10, arr)
	fmt.Println(result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.Trim(string(str), "\r\n")
}
