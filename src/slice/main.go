package main

import "fmt"

func main(){

	a := make([]string, 2, 5)

	a = append(a, "a")

	fmt.Println( a, a[:])



}
