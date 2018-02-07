package demo

import (
	"fmt"
)

type A struct{
	Name string

}

/*
结构的嵌入
 */
type C struct{
	A
	Name string
}

func main(){


	c := C{
		Name: "ys",

		A: A{ Name: "dcl"},
	}

	/*
	当最外层没有该属性时 会去内嵌结构中去找同名结构
	相同级别的结构结构中不能有同名属性
	 */
	c.Name = "yyy"
	c.A.Name = "ssss"

	fmt.Println( b )

}

