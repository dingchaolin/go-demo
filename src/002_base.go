//当前程序的包名
package main

//导入其他的包
//可以取别名
//省略调用
import(
	std "fmt"
	. "fmt"
	myTime "time"
)


//常量的定义
const PI = 3.14

//全局变量的声明与赋值
var name = "dingchaolin"

//结构的定义
type gopher struct {

}

//接口的定义
type golang interface{

}

//由main函数作为函数的入口点启动
func main(){
	std.Println("你好")
	std.Println(myTime.Now())
	Println("你好")

}

