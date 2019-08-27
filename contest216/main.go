package main

import (
	"fmt"
	"math"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//
	//tempS, _, _ := reader.ReadLine()
	//
	//tempArr := strings.Split(string(tempS), " ")
	//
	//a, _ := strconv.ParseInt(tempArr[0], 10, 64)
	//b, _ := strconv.ParseInt(tempArr[1], 10, 64)
	//c, _ := strconv.ParseInt(tempArr[2], 10, 64)
	//
	//tempQ, _, _ := reader.ReadLine()
	//q, _ := strconv.ParseInt(string(tempQ), 10, 64)

	var a, b, c, n int64 = 2, 0, 1, 20

	var maxN int64
	var ind int64

	maxN = a*n*n + b*n + c

	fmt.Println(maxN)

	noPrimes := make([]bool, maxN+1)
	maxDel := int64(math.Floor(math.Sqrt(float64(maxN))))
	noPrimes[0] = true
	noPrimes[1] = true
	for i := int64(2); i <= maxDel; i++ {
		ind = 2 * i
		for ind <= maxN {
			noPrimes[ind] = true
			ind += i
		}
	}
	fmt.Println(noPrimes, len(noPrimes))

	queries := []int64{20}
	queriesMap := map[int64]int64{20: 0}

	//sort.Slice(queries, func(i, j int) bool {
	//	return queries[i] < queries[j]
	//})

	//values := make([]int64, n+1)

	var value int64
	var countPrimes int64

	for i := int64(0); i <= n; i++ {
		value = a*i*i + b*i + c

		if value > 1 && !noPrimes[value] {
			fmt.Println(value, !noPrimes[value])
			countPrimes++
		}

		if _, ok := queriesMap[i]; ok {
			queriesMap[i] = countPrimes
		}
	}

	fmt.Println(noPrimes)
	fmt.Println(queriesMap)
	fmt.Println(queries)
	//for i := int64(0); i < q; i++ {
	//
	//}

}
