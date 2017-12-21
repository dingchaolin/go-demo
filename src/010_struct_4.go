package main

import (
	"fmt"
)

type human struct{
	Sex int

}

/*
结构的嵌入
 */
type teacher struct{
	human
	Name string
	Age int
}

func main(){


	b := teacher{
		Name: "ys",
		Age: 14,
		human: human{ Sex: 1},
	}
/*
两种方式都可以修改属性值
 */
	b.human.Sex = 0
	b.Sex = 0

	fmt.Println( b )

}

