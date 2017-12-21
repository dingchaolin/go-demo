package main

import (
	"fmt"
	"reflect"
)

type User1 struct{
	Id int
	Name string
	Age int

}

type Manager struct{
	User1
	title string
}


func main(){
	m := Manager{ User1:User1{1, "ok", 12}, title:"CTO"}
	t := reflect.TypeOf( m )

	fmt.Printf("%#v\n", t.Field(0)) //取到user1
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0}))//取到user.id
}

