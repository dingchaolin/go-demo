package demo

import (
	"fmt"
)

func main(){

	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println( s1, s2 );

	s1[0] = 111 //都会改变
	fmt.Println( s1, s2 );

	s2 = append( s2, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1)

	s1[0] = 123 //不会影响s2 s2已经分配了新的内存地址
	fmt.Println( s1, s2 );

}