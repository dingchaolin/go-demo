package demo

import (
	"fmt"
	"time"
)

/*
通道实际上是类型化消息的队列
 */

func main(){
	ch := make( chan string )

	go sendData( ch )
	go getData( ch )

	time.Sleep( 2 * 1e9 )


}

func sendData( ch chan string ){
	ch <- "China"
	ch <- "I"
	ch <- "Love"
	ch <- "You"
}

func getData( ch chan string ){
	var input string

	for{
		input = <- ch
		fmt.Println( input )
	}
}