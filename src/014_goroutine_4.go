package main

import (
	"fmt"
	"runtime"
	"sync"//同步包
)

func main(){
	runtime.GOMAXPROCS( runtime.NumCPU() )//设置使用cpu的核数
	wg := sync.WaitGroup{}
	wg.Add( 10 )
	for i := 0; i < 10; i ++{
		go GO1( &wg, i )
	}
	wg.Wait()
}

func GO1( wg * sync.WaitGroup, index int){
	a := 1
	for i := 1; i < 1000000000; i ++ {
		a += i
	}
	fmt.Println( index, a )

	/*
	多核的时候 不能确定 哪个goruntine先结束
	 */
	wg.Done()
}
