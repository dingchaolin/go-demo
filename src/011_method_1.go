package main

import (
	"fmt"
)

type  G struct{
	Name string

}

type H struct{
	Name string
}

func main(){
	g := G{}
	g.Print()

	h := H{}
	h.Print()
}

/*
( g G ) 接受者
Print 方法名称
Print 和 结构 G就连接在了一起
 */
func( g G )Print(){
	fmt.Println( "G" )
}
/*
绑定的类型不同 方法名可以相同  同一类型不能有相同的绑定方法名
参数遵循正常传参规则 值传递 指针传递
 */
func( h H )Print(){
	fmt.Println( "H" )
}