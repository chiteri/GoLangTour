package main 

import "fmt" 

// fibonacci is a function that returns 
// a function that returns an int 
func fibonacci() func(int) int {
	return func(number int) int { 
		a, b := 0, 1

		for i := 0; i < number; i++ {
			a, b = b, a+b
		}

		return b
	}
}

func main () { 
	f := fibonacci() 
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
