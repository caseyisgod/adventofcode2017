package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/caseyisgod/adventofcode2017/routine"
)

type carette struct {
	pos int
}

func main() {
	var car carette
	iteration := car.move(readArrayOfIntFromFile("input.txt"))

	for {
		if i := iteration(); i != 0 {
			fmt.Println(i, "steps")
			break
		}
	}
}
func (car *carette) move(way []int) func() int {
	n, temp := 0, 0
	return func() int {
		n++
		temp = way[car.pos]
		if way[car.pos] >= 3 {
			way[car.pos]--
		} else {
			way[car.pos]++
		}

		car.pos += temp
		if car.pos < 0 || car.pos >= len(way) {
			return n
		}
		return 0
	}
}
func readArrayOfIntFromFile(fn string) []int {
	file, err := os.Open(fn)
	defer file.Close()
	routine.Check_err(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var arr []int
	temp := 0
	for scanner.Scan() {
		temp, err = strconv.Atoi(scanner.Text())
		routine.Check_err(err)
		arr = append(arr, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return arr
}
