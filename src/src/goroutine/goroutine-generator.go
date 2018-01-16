package main

import (
	"math/rand"
	"fmt"
)


func rand_generator()chan int{
	out := make(chan int)
	go func(){
		for{
			out <- rand.Int()
		}
	}()
	return out
}

func main(){
	rand_service_handler := rand_generator()

	fmt.Println( <- rand_service_handler)
}