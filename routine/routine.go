package routine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"fmt"
)

func Check_err(e error) {
	if e != nil {
		panic(e)
	}
}
func GetIndex(i int, leng int) int {
	if i > leng-1 {
		return i - leng
	}
	return i
}
func ReadArrayOfIntFromFile(fn string) []int {
	file, err := os.Open(fn)
	defer file.Close()
	Check_err(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var arr []int
	temp := 0
	for scanner.Scan() {
		temp, err = strconv.Atoi(scanner.Text())
		Check_err(err)
		arr = append(arr, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return arr
}
