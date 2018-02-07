package demo

import (
	"fmt"
)

func generate ( ch chan int ){
	for i:= 2;; i++{
		if i == 100 {
			break;
		}
		fmt.Printf("generate=====:%d\n", i )
		ch <- i

	}
}

func filter( in, out chan int, prime int){
	for{
		i := <- in
		fmt.Printf("filter=====:%d\n", i )
		ret := i % prime
		fmt.Printf("=========:%d,%d,%d\n", i,prime, ret)
		if ret != 0{
			out <- i
		}
	}
}

func main(){
	ch := make( chan int )

	go generate( ch )



	for{

		prime := <- ch
		fmt.Printf("prime=====%d\n", prime)
		ch1 := make( chan int )
		go filter( ch, ch1, prime )
		ch = ch1
	}


}

