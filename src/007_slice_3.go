package main

import (
	"fmt"
)

func main(){

	s1 := make( []int, 3, 6)//default  len == cap
	fmt.Printf( "%p", s1 )
	//如果最终的长度未超过slice的容量则范湖原始slice，将会重新分配数组并拷贝数据
	s1 = append( s1, 4,5,6,7)
    fmt.Printf( "%v %p", s1, s1 )


}