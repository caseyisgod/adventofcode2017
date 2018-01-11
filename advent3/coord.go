package main

import "strconv"

type coord struct {
	X int
	Y int
}

func (c coord) toString() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}
func (c coord) itemsAround(arr *[]string) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			c.X += i
			c.Y += j
			*arr = append(*arr, c.toString())
			c.X -= i
			c.Y -= j
		}
	}
}
