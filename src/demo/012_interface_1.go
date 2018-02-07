package demo

import (
	"fmt"
)

/*
可以所有的结构都实现了empty接口
 */
type empty interface {

}
type  USB interface{
	Name() string
	Connecter//嵌入接口
}

type Connecter interface{
	Connect()
}

type PhoneConnecter struct{
	name string
}

func ( pc PhoneConnecter ) Name() string{
	return pc.name
}

func ( pc PhoneConnecter ) Connect() {
	fmt.Println("connect:", pc.name )
}

func main(){

	var a USB
	a = PhoneConnecter{ name: "PhoneConnecter"}
	a.Connect()
	Disconnect( a )



}


func Disconnect( usb USB ){
	if pc, ok := usb.(PhoneConnecter); ok{
		fmt.Println("Disconnected.", pc.name)
		return
	}
	fmt.Println("Unknown device")
}
