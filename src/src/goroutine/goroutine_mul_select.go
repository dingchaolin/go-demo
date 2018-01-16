package main

import (
	"math/rand"
	"fmt"
)

func rand_generator_2()chan int{
	out := make(chan int)
	go func(){
		for{
			out <- rand.Int()
		}
	}()
	return out
}
/*
输出能力增加1倍
 */
func rand_service() chan int{
	_ran_generator := rand_generator_2()
	_ran_generator2 := rand_generator_2()

	out := make(chan int)

	go func(){
		for{
			out <- <-_ran_generator
		}
	}()

	go func(){
		for{
			out <- <-_ran_generator2
		}
	}()

	return out
}


func main(){
	rand_service_handler := rand_service()

	fmt.Println( <- rand_service_handler)
}