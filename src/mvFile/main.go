package main

import "log"

func test()(a, b int){
	a = 1
	b = 2
	return
}

func main(){
	log.Println(test())
}
