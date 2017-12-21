package main

import (
	"fmt"
)

type Woman struct{
	string
	int

}


func main(){
	/*
	如果是匿名字段 赋值顺序跟定义顺序一致
	 */
	a := person{ "yyy", 13 }
	fmt.Println( a )

}

