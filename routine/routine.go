package routine
import (
	//"fmt"

)
func Check_err(e error) {
	if e != nil {
		panic(e)
	}
}
func GetIndex(i int, leng int) int {
    if i > leng-1 {
        return i-leng
    }
    return i
}