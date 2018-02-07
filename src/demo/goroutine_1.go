package demo

import (
	"fmt"
	"time"
)
/*
如果使用协程 整个代码执行完毕 耗时 10秒
如果去掉go 使用线程 则整个代码执行完毕耗时 17秒
 */
func main(){

	fmt.Println("in main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9 )//纳秒 1e9 表示 1*10^9
	fmt.Println("at the end of main()")

}

func longWait(){
	fmt.Println("beginning longWait()")
	time.Sleep( 5 * 1e9 )// sleep 5s
	fmt.Println("end of longWait()")
}

func shortWait(){
	fmt.Println("beginning shortWait()")
	time.Sleep( 2 * 1e9 )//sleep 2s
	fmt.Println("end of shortWait()")
}