package main 

import ( 
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for count := 0; count < 10; count++ {
		z = z - (( math.Pow(z, 2) - x) / 2 * z) 
		fmt.Println(z)
	} 

	return z
}

func main () {
	fmt.Println(Sqrt(9))
}
