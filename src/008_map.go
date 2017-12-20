package main

import (
	"fmt"
)

func main(){

	var m map[int]string
	m = make(map[int]string)

	var m1 map[int]string = make(map[int]string)

	m2 := make(map[int]string)

	m[1] = "ok"
	a := m[1]
	b := m[2]
	delete( m, 1)
	fmt.Println( m, m1, m2, a, b);

}