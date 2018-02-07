package demo

import (
	"fmt"
)

var ch chan string
func Pingpong( ){
	i := 0
	for{
		fmt.Println( <-ch)
		ch <- fmt.Sprintf("from Pingpong: hi #%d", i )
		i ++
	}
}
func main(){
	ch = make( chan string )

	go Pingpong()

	for i := 0; i < 10; i ++{
		ch <- fmt.Sprintf("from main hello #%d", i)
		fmt.Println( <-ch )
	}



}


