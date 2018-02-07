package demo

import (
	"fmt"
)

const a int = 1
const b = 'A'
const (
	c = a
	d = a + 1
)

//g h i 三个的值相同
//常量必须使用常量进行赋值
// var a = "ddd" b = len(a)  不可以
const (
	g = 1
	h = len("123")
	m = "123"
	u = len( m )
	i
)

// var a = "ddd" b = len(a)  不可以
const (
	gg, hh = 1,2
	// hh  不行 下一行的变量数量必须和上一行一样
	mm, dd // 1, 2

)
const e,r,t = 1,2,4

//枚举类型
// iota  必须要在常量组中使用 从0开始 每定义个常量就自增1
// 常量的iota值和定义的顺序有关系 很 出现的次数没有关系
// 在每一个const组中iota都会从0开始计算
// 常量的定义规则 _MAX_COUNT cMAX_COUNT 这样比较优雅
const(
	aa = ""
	bb
	cc = iota
	ddd
	eee
)

//常量达到枚举的效果
const(
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

/*
计算机存储单位的枚举
 */
const(
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)


 func main(){

	 //fmt.Println( a, b, c, d, e, r, t, g, h, i, mm, dd )
	 fmt.Println( aa,bb,cc,ddd )//65 65 2 3


 }