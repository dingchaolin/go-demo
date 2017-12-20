package main

import (
	"fmt"
)

func main(){
	/*
	defer 传参
	 */
	fmt.Println( "a")
	defer fmt.Println( "b")
	defer fmt.Println( "c") // a c b defer 是逆序调用

	for i := 0; i < 3; i ++{
		defer fmt.Println( i )
	}

}

