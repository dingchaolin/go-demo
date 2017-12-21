package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS( runtime.NumCPU() )//设置使用cpu的核数
	c := make( chan bool, 10 )
	for i := 0; i < 10; i ++{
		go GO( c, i )
	}

	for i := 0; i < 10; i++ {
		<-c //取10次 解决该问题
	}
}

func GO( c chan bool, index int){
	a := 1
	for i := 1; i < 1000000000; i ++ {
		a += i
	}
	fmt.Println( index, a )

	/*
	多核的时候 不能确定 哪个goruntine先结束
	 */
	//if index == 9 {
	//	c <- true
	//}
	c <- true
}
