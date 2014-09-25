package main 

import "fmt" 

type Vertex struct {
	X int 
	Y int 
}

func main () {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	
	// Pointer arithmetic 
	p := Vertex{1, 2}
	q := &p 
	q.X = 1e9
	fmt.Println(p)
}
