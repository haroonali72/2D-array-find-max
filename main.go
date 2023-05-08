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
	arr := make([]int64, n)
	for _, q := range queries {
		a := q[0] - 1
		b := q[1] - 1
		k := int64(q[2])
		arr[a] += k
		if b+1 < n {
			arr[b+1] -= k
		}
	}

	max := int64(0)
	sum := int64(0)
	for _, val := range arr {
		sum += val
		if sum > max {
			max = sum
		}
	}

	return max
}
