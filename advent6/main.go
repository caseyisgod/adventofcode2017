package main

import (
	"fmt"
	"strconv"

	"github.com/caseyisgod/adventofcode2017/routine"
)

type mem []int

func main() {
	var arr mem
	arr = routine.ReadArrayOfIntFromFile("input.txt")

	fmt.Print(arr.redistributionPeriod())
}
func (arr mem) tostring() string {
	temp := ""
	for _, b := range arr {
		temp += strconv.Itoa(b)
	}
	return temp
}

func (arr mem) redistributionPeriod() int {
	m := make(map[string]int)
	m[arr.tostring()] = 0
	for n := 1; true; n++ {
		arr.rearrange()
		if 0 != m[arr.tostring()] {
			return n - m[arr.tostring()]
		}

		m[arr.tostring()] = n
	}
	return 0
}
func (arr mem) rearrange() {
	max, max_index := arr[0], 0

	for i, m := range arr {
		if max < m {
			max = m
			max_index = i
		}
	}
	arr[max_index] = 0
	for i, loci := 1, max_index; i <= max; i++ {
		loci++
		if loci == len(arr) {
			loci = 0
		}
		arr[loci]++
	}
}
