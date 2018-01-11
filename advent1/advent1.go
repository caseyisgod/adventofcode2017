package main
import (
	"fmt"
	"io/ioutil"
	"strconv"
//	"reflect"
//	"os"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//f, err := os.Open("input1.txt")
 	dat, err := ioutil.ReadFile("input.txt")
    check(err)
    arr := make([]int, len(dat)+1)
    for i:=range(dat){
    	arr[i],err =  strconv.Atoi(string(dat[i]))
    	check(err)
    }
    arr[len(dat)] = arr[0]

    var sum int = 0
    for i := 0; i < len(arr)-1; i++ {
		if arr[i]==arr[i+1] {
			sum += arr[i]	
		}
    }

    fmt.Println(sum)
}