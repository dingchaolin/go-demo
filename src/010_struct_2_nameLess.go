package main

import (
	"fmt"
)

type man struct{
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}


func main(){

	/*
	匿名结构
	 */
	b := struct {
		Name string
		Age int
	}{
		Name: "ys",
		Age: 14,
	}


	b.Name = "yyy";

	c := &struct {
		Name string
		Age int
	}{
		Name: "ys",
		Age: 14,
	}

	d := man{ Name:"ccc", Age: 19}
	d.Contact.Phone = "123"//匿名结构赋值
	d.Contact.City = "beijing"

	fmt.Println( b, c.Age, d )

}

