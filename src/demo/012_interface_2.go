package demo

import (
	"fmt"
)

/*
可以所有的结构都实现了empty接口
 */
type empty2 interface {

}
/*
USB2 可以转换为 Connecter2
Connecter2 不可以转换为 USB2

 */
type  USB2 interface{
	Name() string
	Connecter2//嵌入接口
}

type Connecter2 interface{
	Connect()
}

type PhoneConnecter2 struct{
	name string
}

type TVConnecter2 struct{
	name string
}

func ( tv TVConnecter2 ) Connect() {
	fmt.Println("connect:", tv.name )
}

func ( pc PhoneConnecter2 ) Name() string{
	return pc.name
}

func ( pc PhoneConnecter2 ) Connect() {
	fmt.Println("connect:", pc.name )
}

func main(){

	a := PhoneConnecter2{ name: "PhoneConnecter"}
	a.Connect()
	Disconnect2( a )

	tv := TVConnecter2{ name: "TVConnecter2"}
	tv.Connect()
	Disconnect2( tv )

	//此处OK
	var b Connecter2
	b = Connecter2( a )
	b.Connect()
	Disconnect2( b )

	//此处不行 接口没有完全实现
	//var c USB2
	//c = USB( tv )
	//c.Connect()
	//Disconnect2( c )




}

/*
空接口 可以传递很多类型
任何结构都实现了空接口
 */
func Disconnect2( usb interface{} ){


	switch v := usb.(type){
	case PhoneConnecter2:
		fmt.Println("Disconnected.", v.name)
	default:
		fmt.Println("Unknown device")
	}

}
