package main

import "fmt"

/*
and or 多个逻辑运算符的时候 只要最外层不加括号就可以
 */
func main(){
	a := "aa"
	b := "bb"
	if (a == "aa" && len(a) > 0) || (b == "bb" && len(b) > 0) {
		fmt.Println( "正确！")
	}else{
		fmt.Println( "错误!")
	}
}