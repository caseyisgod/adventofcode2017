package main 
import (
	"fmt"
	"os"
	"bufio"
	"github.com/caseyisgod/routine"
	"io"
	"strconv"
//	"reflect"
	)
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	routine.Check_err(err)

	read:=bufio.NewReader(file)
//	fmt.Println(reflect.TypeOf(read))
	var arr [][]int
	fill_array(&arr,read)
	fmt.Println(CheckSum(&arr))
}
func CheckSum(a *[][]int) int {
	sum, min, max := 0,0,0
	for _,temp := range (*a){
		min, max = temp[0],temp[0]
		for _,e :=range (temp){
			if e < min{min=e}
			if e > max{max=e}
		}
		sum += max - min
	}
	return sum
}
func fill_array(arr *[][]int,read *bufio.Reader) {
	for i:=0; true; i++ {
	var temp []int
	status := "ok"
	buf := ""
	value:=0
	
	for j:=0; true; j++ {

		buf = ""
		status = "ok"
		for status == "ok" {      //reads bytes, converts them to string and append to existing arr[i][j] element until \n or \t
			//after avery line we have /r carriage return symbol and then \n new line symbol 
			token, err := read.ReadByte()

			if err == io.EOF {
				status = "EOF"
				value,err = strconv.Atoi(buf)
				temp = append(temp,value)
				routine.Check_err(err)
				fmt.Printf("%v\n",temp[j])
				break
				} else {routine.Check_err(err)}

			if (token != byte('\t'))&&(token != byte('\r')) {
				buf += string(token)
				
				} else if token == byte('\t') {
				status = "tab"
				value,err = strconv.Atoi(buf)
				temp = append(temp,value)
				routine.Check_err(err)
				fmt.Printf("%v\t",temp[j])	
				break

				} else {status = "newline"
				value,err = strconv.Atoi(buf)
				temp = append(temp,value)
				routine.Check_err(err)
				discard, err := read.Discard(1) //discarding \n symbol after \r
				if discard !=1{routine.Check_err(err)}
				fmt.Printf("%v\n",temp[j])	
				
				break}
			}
		if (status == "newline")||(status == "EOF") {
			break
			}
		}
		*arr = append(*arr,temp)
	
	if status == "EOF" {
			break
		}
	}
}