package main

import (
	"fmt"
)

type  TZ int

func main(){


	var tz TZ
	tz.Print()// method value
	(*TZ).Print(&tz)//method expression
	tz.Increase( 100 )
	fmt.Println( tz )
}

/*
方法名称冲突和字段名称冲突解决方法一样
通过类型绑定 可以对任何类型绑定方法
如果TZ 定义在别的包中 这个方法就不能调用了 因为找不到和哪个方法进行绑定
 */
func( tz TZ )Print(){
	fmt.Println( "TZ" )//可以输出
}

func( tz *TZ )Increase( num int ){
	*tz += TZ(num)
}