package main

import (
	"fmt"
)

func main(){

	c := make( chan bool )

	go func(){
		fmt.Println("GOGOGO")
		c <- true
		close( c )//退出时 for range也会停止  main退出  否则 不关闭  for range会变成死锁
	}()
	for v := range c {
		fmt.Println( v )
	}

}

