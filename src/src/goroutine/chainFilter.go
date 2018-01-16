package main

import (
	"fmt"
)

func generate(ch chan<- int) { /*只写通道*/
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int /*只读通道*/ , out chan<- int /*只写通道*/ , prime int, index int) {
	for {
		i := <-in
		fmt.Println( fmt.Sprintf("i======%d,prime====%d, index===%d", i, prime, index))
		if i%prime != 0 {
			out <- i
		}
	}
}
/*
具体结合channel-filter.png理解
 */
func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 10; i ++ {
		prime := <-ch
		fmt.Println("prime====", prime)
		ch1 := make(chan int)
		go filter( ch, ch1, prime, i )
		ch = ch1
	}


}
