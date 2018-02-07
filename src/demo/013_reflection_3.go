package demo

import (
	"fmt"
	"reflect"
)


/*

想要修改的时候 必须传指针
 */

func main(){
	x := 123
	v := reflect.ValueOf( &x )//必须是指针才能修改
	v.Elem().SetInt( 999 )

	fmt.Println( x )
}

