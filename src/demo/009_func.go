package demo

import (
	"fmt"
)

func main(){
	a, b := 1,2
	funcE("OK", a,b) //里面不会修改到a,b的值 值拷贝
	fmt.Println( a,b )

	s1 := []int{1,2,3,4}
	funcF( s1) //里面会修改到s1的值 地址拷贝

	c := 3
	funcG( &c )
	fmt.Println(  c )

	F := funcG

	d := 5
	F( &d )
	fmt.Println( d )

	/*
	匿名函数
	 */
	G := func( msg string ){
		fmt.Println( msg )
	}

	G( "nihao")

	H := closure( 10 ) //返回一个函数
	fmt.Println( H( 1 ))//11
	fmt.Println( H( 2 ))//12

}

func funcA( a int, b string)(int, string){

	return 1,"ok"
}

func funcB( a int)int{

	return 1
}

func funcC( a,b,c int)(int,int,int){

	return 1,2,3
}

/*
命名返回值参数
命名的返回值在函数体内已经存在 不需要再定义
 */
func funcD( )(a,b,c int){
	a, b, c = 1, 2, 3
	return a, b, c //返回值都要写上 代码可读性更高
}

/*
不定长变参
最后一个单数是不定长变参
 */

 func funcE( b string, s ...int){
 	//此时s是一个slice
 	s[0] = 111
 	s[1] = 222
 	fmt.Println( b, s )
 }

func funcF(  s []int){
	//此时s是一个slice
	s[0] = 111
	s[1] = 222
	fmt.Println(  s )
}

func funcG( s *int){
	*s = 222


}

/*
闭包
地址全相同
 */
func closure( x int ) func(int) int{
	fmt.Printf("closure===%p\n", &x)
 	return func( y int ) int{
 		fmt.Printf("closure return====%p\n", &x)
 		return x + y
	}
 }