package main

import "fmt"

type Person struct{
	Name string
	Age int
	Sex int
}
func main(){
	arr := Person{"aaa", 12, 1}
	for name:= range arr {
		fmt.Println( name )
	}
}
