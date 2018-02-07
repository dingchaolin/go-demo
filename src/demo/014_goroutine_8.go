package demo

import (
	"fmt"
	"time"
)

func main(){
	c := make( chan bool )
	select{
	case v := <- c:
		fmt.Println( v )
		//time.After 返回 <-chan Time
	case <- time.After( 3 * time.Second ):
		fmt.Println("time out")
	}


}


