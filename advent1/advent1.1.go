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
 	dat, err := ioutil.ReadFile("input1.txt")
    check(err)
    arr := make([]int, len(dat))
    for i:=range(dat){
    	arr[i],err =  strconv.Atoi(string(dat[i]))
     	check(err)
    }

    var sum int = 0

    l := len(arr)
    for i := 0; i < l; i++ {

		if arr[i]==arr[getIndex(i+l/2,l)] {
			sum += arr[i]	
		}
    }

    fmt.Println(sum)
}
 func getIndex(i int, leng int) int {
    if i > leng-1 {
        return i-leng
    }
    return i
}