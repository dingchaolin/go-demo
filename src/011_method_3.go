package main

import (
	"fmt"
)

type  Y struct{
	name string
}

func main(){


	 y := Y{}
	y.Print()// method value
	fmt.Println( y.name )
}

/*
方法可以访问私有属性
私有 是指 在整个package 而言的 超出 package 私有属性就不能访问了
 */
func( y *Y )Print(){
	y.name = "qwe"
	fmt.Println( y.name )//可以输出
}