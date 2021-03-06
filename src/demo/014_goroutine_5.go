package demo

import (
	"fmt"
)

func main(){
	c1, c2 := make( chan int), make( chan string)
	o := make( chan bool, 2 )

	go func(){

		for{
			select{
			case v, ok := <- c1:
				if !ok {
					o <- true

					break
				}
				fmt.Println( "c1", v )
			case v, ok := <- c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println( "c2", v )
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	/*
	如果关闭了其中一个，另一个就会一直返回0 ok=false 状态 程序无法正常退出
	所以 应该是 只要关闭一个 程序就应该退出
	同时关闭两个是不可行的
	 */
	close( c1 )
	//close( c2 )

	for i := 0; i < 2 ; i++{
		<- o
	}

}


