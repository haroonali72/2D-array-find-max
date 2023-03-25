package main

import "fmt"

func main() {
	fmt.Println(arrayManipulation(10, [][]int32{
		[]int32{1, 5, 3},
		[]int32{4, 8, 7},
		[]int32{6, 9, 1},
	}))
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	manipulatedData := zeroValues2DArray(n, queries)
	max := int64(0)

	for i, _ := range manipulatedData {
		var j int32
		for j = 0; j < n; j++ {
			if j+1 >= queries[i][0] && j+1 <= queries[i][1] {
				if i > 0 {
					manipulatedData[i][j] = queries[i][2] + manipulatedData[i-1][j]
				} else {
					manipulatedData[i][j] = queries[i][2]
				}
			} else if i > 0 {
				manipulatedData[i][j] = manipulatedData[i-1][j]
			}
		}

		max = compareAndFindMax(max, manipulatedData[i])
	}

	return max
}

func zeroValues2DArray(n int32, queries [][]int32) [][]int32 {
	emptyData := make([][]int32, len(queries))
	for index, _ := range emptyData {
		emptyData[index] = make([]int32, n)
	}

	return emptyData
}

func compareAndFindMax(max int64, nums []int32) int64 {
	m := int32(max)

	for _, n := range nums {
		if m < n {
			m = n
		}
	}

	return int64(m)
}
