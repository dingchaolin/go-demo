//当前程序的包名
package demo

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

//组定义
const (
	a = 1
	b = 2
	c = 3
)

//只能声明全局变量 函数体内不能使用
var (
	d = "fff"
	e = ""
	g = 234
)

type(
	newType int
	type1 float32
)


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

