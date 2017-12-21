package main

import (
	"fmt"
	"reflect"
)

type User2 struct{
	Id int
	Name string
	Age int
}


func main(){
	u :=  User2{1, "OK", 12}
	Set( &u )
	fmt.Println( u )

}

func Set( o interface{}){
	// v.Elem() 返回的是 reflect.ValueOf
	v := reflect.ValueOf( o )

	if v.Kind() != reflect.Ptr || !v.Elem().CanSet(){
		fmt.Println("XX")
		return
	}else{
		v = v.Elem()
	}

	f := v.FieldByName("Name");
	if !f.IsValid(){
		fmt.Println("BAD")
		return
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("nihao")
	}
}
