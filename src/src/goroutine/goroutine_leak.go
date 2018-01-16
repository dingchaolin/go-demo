package main

import (
	"time"
	"fmt"
)

func never_leak(ch chan int) {
	timeout := make(chan bool, 1)

	/*
	启动timeout协程
	 */
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- true
	}()

	/*
	监听通道
	 */

	 select {
	 case <- ch:
	 	case <- timeout:
	 		fmt.Println("timeout")
	 		return
	 }
}

/*
使用超时避免读堵塞，使用缓冲避免写堵塞。
 */
func main(){
	ch := make(chan int)
	never_leak( ch )
}