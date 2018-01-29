package main

import (
	"yaml/config"
	"fmt"
)

func main(){
	conf, err := conf.NewConfig1("./server.yaml")
	if err != nil{
		panic( err )
	}
	fmt.Println( conf )
}