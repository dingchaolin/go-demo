package demo

import (
	"fmt"
)

func main(){

	/*
	如果函数体内某个变量作为defer时函数的参数 则在定义defer时即已经获得了值拷贝，
    如果不是参数，使用时已经在函数体内定义了改参数 则是引用某个变量的地址
	 */
    var fs = [4]func(){}

    for i := 0; i < 4; i ++{
    	/*
    	此处i是作为函数的参数传入defer的，此时获得是参数的值拷贝
    	 */
    	defer fmt.Println("defer i = ", i)
		/*
		此处是闭包，i是main中的一个局部变量传入的，获取的是i的地址
		 */
    	defer func(){
			fmt.Println("defer 闭包 i = ", i)
		}()
		/*
		此处是闭包，i是main中的一个局部变量传入的，获取的是i的地址
		 */
		fs[i] = func(){
			fmt.Println("闭包 i = ", i)
		}
	}

	for _, f := range fs{
		f()
	}

}

