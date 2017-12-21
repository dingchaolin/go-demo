package main

import (
	"fmt"
)

/*
可以所有的结构都实现了empty接口
 */
type empty3 interface {

}
type  USB3 interface{
	Name() string
	Connecter3//嵌入接口
}

type Connecter3 interface{
	Connect()
}

type PhoneConnecter3 struct{
	name string
}

func ( pc PhoneConnecter3 ) Name() string{
	return pc.name
}

func ( pc PhoneConnecter3 ) Connect() {
	fmt.Println("connect:", pc.name )
}

func main(){
/*
接口的方法调用和结构的方法调用时不同的
如果传进的是指针 则可以调用receiver时指针的方法 也 可以调用receiver不是指针的方法
如果传进的非指针 那就不能调用指针的方法集 只能调用属于自己的那部分
 */

	pc := PhoneConnecter3{ name: "PhoneConnecter"}
	var a  Connecter3
	//将对象赋值给接口时 会发生拷贝 而接口内部存储的是指向这个复制品的指针 既无法修改赋值品的状态 也不无获取指针

	a = Connecter3( pc )
	a.Connect();

	pc.name = "pc"//此处无效
	a.Connect()

	var d interface{}
	fmt.Println( d == nil ) //true
	//只有当接口存储的类型和对象都为nil时 接口才等于nil
	var p *int = nil
	d = p
	fmt.Println( d == nil ) //false



}


func Disconnect3( usb USB3 ){
	if pc, ok := usb.(PhoneConnecter3); ok{
		fmt.Println("Disconnected.", pc.name)
		return
	}
	fmt.Println("Unknown device")
}
