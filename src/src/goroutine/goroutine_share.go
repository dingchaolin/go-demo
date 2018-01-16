package main

import "fmt"

type share struct {
	reader chan int
	writer chan int
}

func watchdog(v share) {

	go func() {
		var value = 0

		for {
			select {
			case value = <-v.writer:
			case v.reader <- value:

			}
		}
	}()
}


func main(){
	// 对于协程安全的共享变量 也不用使用锁来防止资源读写冲突
	v := share{ make(chan int), make(chan int)}
	watchdog( v )

	fmt.Println( <- v.reader)
	v.writer <- 666
	fmt.Println( <- v.reader )

}
