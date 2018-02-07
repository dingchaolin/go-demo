package demo

import (
	"fmt"
	"reflect"
)

type User3 struct{
	Id int
	Name string
	Age int
}

func (u User3 ) Hello( name string, name2 string) string{
	fmt.Println("hello", name, name2, "my name is ", u.Name )
	return name2
}

func main(){
	u :=  User3{1, "dcl", 12}

	//u.Hello("ys")

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("Jay"),reflect.ValueOf("JJ")}
	ret := mv.Call( args )
	fmt.Println( ret )
}

