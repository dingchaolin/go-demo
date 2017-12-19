package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
零值
 */

 type(
 	byte int8
 	rune int32
 	字符串 string
 )

 func main(){
 	//var a int32 //0
	// var a bool //false
	// var a string //""
	//var a []int //[] slice
	// var a [5]int //[0,0,0,0,0]
	 var a [5]bool //[false false false false false]

 	fmt.Println( a )
 	fmt.Println( math.MaxInt8 )//127 可以使用该值对变量检查是否溢出

 	var str 字符串
 	str = "nihao"
 	fmt.Println( str )


 	var b int = 1
 	fmt.Println( b )

 	c := 2
 	fmt.Println( c )

 	var e, f, g int
 	e, f, g = 1, 2, 3
	 fmt.Println( e, f, g )

	 var h, j int = 4,5
	 fmt.Println( h,j)

	 k,l := 6,7
	 fmt.Println( k,l)

	 //空白符号 对赋值进行忽略
	 //用于函数中忽略返回值
	 n,_,m:=8,9,10
	 fmt.Println( n, m )

	 //变量的类型转换 可以转换 会造成精度丢失
	 var t float32 = 111165.1; y:= int(t)
	 fmt.Println( t, y )

	 v:= string( y )
	 fmt.Println( v )//转成ascii字符

	 vStr := strconv.Itoa( y )
	 fmt.Println( vStr )
	 vInt,_ := strconv.Atoi( vStr )
	 fmt.Println( vInt )









 }