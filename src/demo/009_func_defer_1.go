package demo

import (
	"fmt"
)

func main(){

    //defer 局部变量
	for i := 0; i < 3; i ++{

		defer func(){
			fmt.Println( i )//3 3 3
		}()
	}

	//main 退出的时候开始执行 defer
	// defer 内部引用的是i这个局部变量  退出时 i=3 所以 打印的都是3

}

