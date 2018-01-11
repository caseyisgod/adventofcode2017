package main

//taxicad distance calculation on spiral numbers distribution
import (
	"fmt"
)

func main() {
	im := new(int)
	fmt.Scanf("%d", im)
	var number = *im
	/*define n as circle number
	number of places in a circle is n*8. That means series is 1 + 8*n
	last number in the nth circle is a sum of first n terms in arithmetic progression 4n^2 - 4n + 1*/
	var position coord
	position.X, position.Y = positionOfTheItem(number)
	/*Here starst the second part of the task.
	We have a function that shows coordinades of n'th memory unit relative to the first item wich is 1 and located at 0,0
	Now the map will be created, that holds all the values in previously filled memory units.
	Then it will be easy to find all adjasent units and find out value in the next one*/

	a, b := fillMemoryUpToVanue(number)
	fmt.Print(a, ", position: "+b)
} //

func fillMemoryUpToVanue(val int) (int, string) {
	//creating map
	m := make(map[string]int)
	//determine first element
	m["0,0"] = 1
	//creating temporary position variable
	var (
		position coord
		temp     int
	)
	//this array contains surrounding items' coordinates
	var arr []string
	//iterating over spiral
	for k := 2; true; k++ {
		position.X, position.Y = positionOfTheItem(k)
		arr = nil
		temp = 0
		position.itemsAround(&arr)
		for _, i := range arr {

			temp += m[i]
		}
		m[position.toString()] = temp
		if temp > val {
			break
		}
	}
	return temp, position.toString()
}

/*func fillMemoryUpTo(n int) lastItem int {
	//creating 4 sides of the square for both current and previous circles (lol)
	prev, curr := make([][]int,4),make([][]int,4)
	status := "missing"
	for i := range curr {
		append(curr[i],1) //filling first element
	}
	for i := 1; "found"!=status; i++ { //running throw cicles
		prev = curr
		for j := 0; j < 4; j++ { ////running throw sides
			for k := 0; k < i*2; k++ { //running throw slots on side
				switch  {
				case i*2-1==k:

				case i*2-2==k:

				case 0==j&&0==k: //if we're fillig first element we sum last and first in preveous cicle
					curr[j][k]=prev[0][0]+prev[3][(i-1)*2-1]
				case 3==j&&i*2-1==k: //if we're fillig last element we sum own first, own previous and last in preveous cicle
					curr[j][k]=prev[3][(i-1)*2-1]+curr[0][0]+curr[j][k-1]
				default:

				}
			}
		}
	}
}*/

func positionOfTheItem(number int) (int, int) {
	//first, we search the n of imputed number's circle
	nc := determineCicle(number)
	//second, we find the side the number lies on
	//distance := 0
	x, y := 0, 0
	for i := 1; i < 5; i++ {
		if number <= maxInCicle(nc-1)+i*nc*2 {
			//third, we find out how far the number is from the side's center
			lengthOnSide := number - (maxInCicle(nc-1) + i*nc*2 - nc) /*side's center*/
			//and, the last, we find the distance to the center of the spiral
			//distance = int(math.Abs(float64(lengthOnSide))) + nc
			switch i {
			case 1:
				x = nc
				y = lengthOnSide
			case 2:
				x = -lengthOnSide
				y = nc
			case 3:
				x = -nc
				y = -lengthOnSide
			case 4:
				x = lengthOnSide
				y = -nc
			}
			break
		}
	}
	//return "distance is " + strconv.Itoa(distance) + ", x is " + strconv.Itoa(x) + ", y is " + strconv.Itoa(y)
	return x, y
}
func determineCicle(num int) int {
	for i := 1; true; i++ {
		if num <= maxInCicle(i) {
			return i
		}
	}
	return 0
}
func maxInCicle(n int) int {
	i := n + 1
	return 4*i*(i-1) + 1
}
