package demo

import (
	"fmt"
)

func main(){

	//copy 的slice之间 不会相互影响
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9}
	//copy( s1, s2 )//s1 [7 8 9 4 5 6] s2 [7 8 9]
	copy( s2, s1 )//s1 [1 2 3 4 5 6] s2 [1 2 3] //拷贝以短的为准
	copy( s2, s1[1:3] ) //拷贝局部
	copy( s2[0:2], s1[1:3] ) //拷贝局部 到指定位置
	fmt.Println( s1, s2 );

	s3:= s1[0:5] //完成一个整的拷贝
	s4:= s1[:] //完成一个整的拷贝

	fmt.Println( s3, s4 );

	test( s1 )

	fmt.Println( s1 );

	sa := []int{1,2,3,4,5,6}
	sa = append( sa[:5], sa[6:]...)
	fmt.Println( sa );




}

func test( s []int ) []int {
	s[0] = 555
	return s
}