package main

import (
	"fmt"
	"time"
)

/*
通道实际上是类型化消息的队列
 */
 
func main(){
	ch := make( chan int )

	go pump( ch )

	go suck( ch )


	time.Sleep( 10 * 1e9 )
}

func pump( ch chan int ){
	for i := 0;;i++{
		ch <- i;
	}
}

func suck( ch chan int )  {
	for i := 0;; i++{
		fmt.Println( <- ch )
	}
}
